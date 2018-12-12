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
