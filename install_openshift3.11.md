network manager会调用dnsmasq，请确保调用前/etc/resolv.conf内容是正确的，比如内容是：nameserver 114.114.114.114 在调用后会在/etc/dnsmasq.d中生成origin-dns.conf 和 origin-upstream-dns.conf，大致原理是：dnsmasq会比较resolv.conf的内容和当前的内容是否一致（MD5）不一致会生成origin-upstream-dns.conf，主要逻辑在/etc/NetworkManager/dispatcher.d/99-origin-dns.sh中
## 上述那个脚本会检查有没有99-origin-dns.sh，如果有，不会生成upstream 所以，要从resolv.conf中删掉这一行
今天解决了svc无法解析的问题：svc在openshift的里的流程是：这个稍有些复杂，首先它由/etc/resovle.conf、dnsmasq、dbus、openshift进程组成。
resovle.conf转移请求到node也就是本机的53端口，这个端口由dnsmasq监听，它的配置文件是/etc/dnsmasq.d/origin-dns.conf，其中有几行需要注意：
except-interface=lo是监听除了lo外的所有网卡，所以用netstat会看到dnsmasq监听几个interface。
enable-dbus 这个一定要有，它是让openshift更新upstream的入口，也就是说除了dnsmasq中的配置upstream（/etc/dnsmasq.d/origin-upstream-dns.conf），openshift还会写两个upsteam，systemct status dnsmasq可以看到如下输出
   using nameserver 127.0.0.1#53 for domain in-addr.arpa
   using nameserver 127.0.0.1#53 for domain cluster.local
以上两个会把请求转发到openshift监听的本机53端口上，完成域名到ip地址之间的转换，然后可以通过本地iptables实现流量的转发（可以通过iptable-save查看转发规则）
刚才说到的往dbus写upsteam的进程其实是一个pod，名字是sdn-xxx，可以用 oc logs sdn-xxx -n openshift-sdn查看log，源码在https://github.com/openshift/origin/blob/e004e6513ec755a7a106dc97516703847e470d0b/pkg/dns/dnsmasq.go中

cat /etc/sysconfig/network-scripts/ifcfg-eth0 
NM_CONTROLLED=yes

yum -y install wget git net-tools bind-utils yum-utils iptables-services bridge-utils bash-completion kexec-tools sos psacct

yum -y install  https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm 

sed -i -e "s/^enabled=1/enabled=0/" /etc/yum.repos.d/epel.repo 

yum -y --enablerepo=epel install ansible pyOpenSSL 

yum install -y docker-1.13.1 

systemctl enable docker 

systemctl start docker 

yum install origin origin-docker-excluder origin-excluder -y 

config ssh  

install ansible 

config okd_hosts 
###################################################################################################################
# Create an OSEv3 group that contains the master, nodes, etcd, and lb groups.
# The lb group lets Ansible configure HAProxy as the load balancing solution.
# Comment lb out if your load balancer is pre-configured.
[OSEv3:children]
masters
nodes
etcd

# Set variables common for all OSEv3 hosts
[OSEv3:vars]
ansible_ssh_user=root
openshift_deployment_type=origin
openshift_disable_check=docker_storage,memory_availability,disk_availability,docker_image_availability,package_availability

openshift_image_tag="v3.11"
# uncomment the following to enable htpasswd authentication; defaults to AllowAllPasswordIdentityProvider
#openshift_master_identity_providers=[{'name': 'htpasswd_auth', 'login': 'true', 'challenge': 'true', 'kind': 'HTPasswdPasswordIdentityProvider'}]
openshift_additional_repos=[{'id': 'centos-okd-ci', 'name': 'centos-okd-ci', 'baseurl' :'https://rpms.svc.ci.openshift.org/openshift-origin-v3.11', 'gpgcheck' :'0', 'enabled' :'1'}]

openshift_master_cluster_method=native
openshift_master_cluster_hostname=baas-test-env

openshift_web_console_nodeselector={'node-role.kubernetes.io/master':'true'}

# apply updated node defaults
#openshift_node_groups=[{'name': 'node-config-all-in-one', 'labels': ['node-role.kubernetes.io/master=true', 'node-role.kubernetes.io/infra=true', 'node-role.kubernetes.io/compute=true'], 'edits': [{ 'key': 'kubeletArguments.pods-per-core','value': ['20']}]}]

# host group for masters
[masters]
okd-master1
okd-master2
okd-master3

# host group for etcd
[etcd]
okd-master1
okd-master2
okd-master3

# host group for nodes, includes region info
[nodes]
okd-master1 openshift_node_group_name='node-config-master-infra'
okd-master2 openshift_node_group_name='node-config-master-infra'
okd-master3 openshift_node_group_name='node-config-master-infra'
okd-node1 openshift_node_group_name='node-config-compute'
okd-node2 openshift_node_group_name='node-config-compute'
okd-node3 openshift_node_group_name='node-config-compute'

###################################################################################################################

scp /etc/yum.repos.d/CentOS-OpenShift-Origin311.repo okd-node4:/etc/yum.repos.d/CentOS-OpenShift-Origin311.repo
yum -y install wget git net-tools bind-utils yum-utils iptables-services bridge-utils bash-completion kexec-tools sos psacct NetworkManager origin origin-docker-excluder origin-excluder origin-node.x86_64
yum -y install  https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sed -i -e "s/^enabled=1/enabled=0/" /etc/yum.repos.d/epel.repo
yum -y --enablerepo=epel install ansible pyOpenSSL
yum install -y docker-1.13.1
systemctl enable docker
systemctl enable NetworkManager
systemctl start NetworkManager 

rm -rf /etc/sysconfig/docker-storage
scp /etc/sysconfig/docker-storage-setup
docker-storage-setup

pvcreate /dev/vdb  创建pv
vgcreate docker-vg /dev/vdb 创建vg
lvcreate docker--vg-docker--pool  创建lv
lvextend -L 190G /dev/docker-vg/docker-pool 扩展
lvconvert --type thin --thinpool 转换成薄模式

systemctl start docker 

echo 1048576 > /proc/sys/fs/inotify/max_user_watches
echo 1 > /proc/sys/net/ipv4/ip_forward

ansible-playbook -i ./okd.hosts ~/openshift-ansible/playbooks/adhoc/uninstall.yml

ansible-playbook -i okd.hosts openshift-ansible/playbooks/prerequisites.yml

ansible-playbook -i ./okd.hosts /root/openshift-ansible/playbooks/deploy_cluster.yml
