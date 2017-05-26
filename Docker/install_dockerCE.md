#https://docs.docker.com/engine/installation/linux/rhel/#install-from-a-package
yum install -y yum-utils
yum-config-manager     --add-repo     https://download.docker.com/linux/centos/docker-ce.repo
yum makecache fast
yum -y install docker-ce
