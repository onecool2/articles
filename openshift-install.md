### Openshift Install

这片文章介绍如何安装Openshift 1.6，现在主流的方式有2两种，一种是快速安装 oc cluster up，比较适合于快速部署一个单机版本的openshift，但是要安装多节点版本这种就
---
### 1 Openshif对操作系统的要求：
Fedora 21, CentOS 7.3, RHEL 7.3, or RHEL 7.4或者RHEL Atomic Host 7.3.6 or later.

### 2 硬件配置：
最少2 core CPU， 最小16 GB RAM. 在/var目录下最少40 GB剩余空间，在/usr/local/bin目录下最少1GB剩余空间，/tmp目录下最少1GB的剩余空间。

### 3 开始执行安装命令：
yum install -y wget git net-tools bind-utils iptables-services bridge-utils bash-completion kexec-tools sos psacct

yum update -y

yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm 

sed -i -e "s/^enabled=1/enabled=0/" /etc/yum.repos.d/epel.repo

yum -y --enablerepo=epel install ansible pyOpenSSL

yum install -y docker-1.12.6

cd ~

git clone https://github.com/openshift/openshift-ansible

cd openshift-ansible

ssh-copy-id host
  docker pull openshift/origin-deployer:v3.6.1 
  docker pull openshift/origin-docker-registry:v3.6.1
  docker pull openshift/origin-haproxy-router:v3.6.1
  docker pull openshift/origin-pod:v3.6.1

下面是ansible中host文件的内容
<pre><code>
# Create an OSEv3 group that contains the masters, nodes, and etcd groups
[OSEv3:children]
masters
nodes
# Set variables common for all OSEv3 hosts
[OSEv3:vars]
ansible_ssh_user=root
openshift_deployment_type=origin
openshift_disable_check=disk_availability,docker_image_availability,docker_storage
# uncomment the following to enable htpasswd authentication; defaults to DenyAllPasswordIdentityProvider
# openshift_master_identity_providers=[{'name': 'htpasswd_auth', 'login': 'true', 'challenge': 'true', 'kind':  'HTPasswdPasswordIdentityProvider', 'filename': '/etc/origin/master/htpasswd'}]
	
[etcd]
master1-26-95
# host group for masters
[masters]
master1-26-95

# host group for nodes, includes region info
[nodes]
master1-26-95 openshift_node_labels="{'region': 'infra', 'zone': 'default'}"
node1-26-97 openshift_node_labels="{'region': 'primary', 'zone': 'east'}"
node2-26-98 openshift_node_labels="{'region': 'primary', 'zone': 'east'}"
node3-26-99 openshift_node_labels="{'region': 'primary', 'zone': 'east'}"
<pre></code>

### 4 然后执行：ansible-playbook ~/openshift-ansible/playbooks/byo/config.yml

disable mutiple NIC.
add master hostname into /etc/hosts
##################################################################################
1. unset https_proxy http_proxy
2. cp centos7.json centos.json
3. a) lvremove /dev/cl/home
   b) echo VG=cl > /etc/sysconfig/docker-storage-setup
   c) docker-storage-setup
   d) lvextend -L +50G /dev/cl/docker-pool
   e) systemctl start docker
   # pv->vg->lv we take the second way to install docker pool
   # vgreduce ?

###############################
export KUBECONFIG=/etc/origin/master/admin.kubeconfig
oc login -u system:admin
oadm policy add-cluster-role-to-user cluster-admin admin
oadm manage-node master1-40-12 --schedulable
docker load < /root/origin-pod.tar;docker tag 77b5b3e452aa openshift/origin-pod:v3.6.1;docker load < /root/origin-deployer.tar;docker tag 90fbedb07cc9 openshift/origin-deployer:v3.6.1
oadm policy add-scc-to-user hostnetwork -z router
oadm router router --replicas=1 --service-account=router

####################harbor
install docker-compose
modify harbor.cfg master1-40-12:5080 
modify docker-compose.yml  80-> 5080
modify /etc/sysconfig/docker insecure-registry=master1-40-12:5080



