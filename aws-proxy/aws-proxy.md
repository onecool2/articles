# On EC2:
```
docker run -d --name centos7 -p3128:3128 --privileged centos:7 /usr/sbin/init 
docker exec -it centos7 /bin/bash

cat << EOF > /etc/yum.repos.d/CentOS-Base.repo
[base]
name=CentOS-$releasever - Base - mirrors.aliyun.com
failovermethod=priority
baseurl=https://mirrors.aliyun.com/centos-vault/7.9.2009/os/$basearch/
gpgcheck=1
gpgkey=https://mirrors.aliyun.com/centos-vault/RPM-GPG-KEY-CentOS-7
 
#released updates 
[updates]
name=CentOS-$releasever - Updates - mirrors.aliyun.com
failovermethod=priority
baseurl=https://mirrors.aliyun.com/centos-vault/7.9.2009/updates/$basearch/
gpgcheck=1
gpgkey=https://mirrors.aliyun.com/centos-vault/RPM-GPG-KEY-CentOS-7
 
#additional packages that may be useful
[extras]
name=CentOS-$releasever - Extras - mirrors.aliyun.com
failovermethod=priority
baseurl=https://mirrors.aliyun.com/centos-vault/7.9.2009/extras/$basearch/
gpgcheck=1
gpgkey=https://mirrors.aliyun.com/centos-vault/RPM-GPG-KEY-CentOS-7
 
#additional packages that extend functionality of existing packages
[centosplus]
name=CentOS-$releasever - Plus - mirrors.aliyun.com
failovermethod=priority
baseurl=https://mirrors.aliyun.com/centos-vault/7.9.2009/centosplus/$basearch/
gpgcheck=1
enabled=0
gpgkey=https://mirrors.aliyun.com/centos-vault/RPM-GPG-KEY-CentOS-7
 
#contrib - packages by Centos Users
[contrib]
name=CentOS-$releasever - Contrib - mirrors.aliyun.com
failovermethod=priority
baseurl=https://mirrors.aliyun.com/centos-vault/7.9.2009/contrib/$basearch/
gpgcheck=1
enabled=0
gpgkey=https://mirrors.aliyun.com/centos-vault/RPM-GPG-KEY-CentOS-7
EOF
```

```
yum install -y squid openssl openssl-devel stunnel httpd
openssl req -new -x509 -days 3650 -nodes -out stunnel.pem -keyout stunnel.pem  
openssl gendh 512>> stunnel.pem
```

```
cat << EOF >> /etc/stunnel/stunnel.conf
cert = /etc/stunnel/stunnel.pem
CAfile = /etc/stunnel/stunnel.pem
socket = l:TCP_NODELAY=1
socket = r:TCP_NODELAY=1
pid = /var/log/stunnel.pid
verify = 3
compression = zlib
delay = no
sslVersion = TLSv1
fips=no
debug = 7
syslog = no
output = /var/log/stunnel.log

[squid]
accept = 443
connect = 127.0.0.1:3128
EOF
```
```
openssl req -new > proxy.csr
openssl rsa -in privkey.pem -out proxy.key
openssl x509 -in proxy.csr -out proxy.crt -req -signkey proxy.key -days 3650
```
* /etc/squid/squid.conf
```
# acl SSL_ports port 80
acl SSL_ports port 443

# Squid normally listens to port 3128
http_port 3128 cert=/etc/squid/cert/proxy.crt key=/etc/squid/cert/proxy.key
```

* if you want to add username and password for your proxy
```
htpasswd -c /etc/squid/user.pass username
in the begin of the squid.conf

auth_param basic program /lib64/squid/basic_ncsa_auth /etc/squid/user.pass
auth_param basic children 5
auth_param basic realm Welcome to test
auth_param basic credentialsttl 2 hours
acl ncsa_users proxy_auth REQUIRED
dns_nameservers 8.8.8.8
http_access allow ncsa_users

via off
forwarded_for delete
```
```
service squid start
/usr/bin/stunnel
```

# On your local server:
* scp stunnel.pem to your local server
```
yum install stunnel
```
```
cat << EOF >>/etc/stunnel/stunnel.conf
cert = /etc/stunnel/stunnel.pem
socket = l:TCP_NODELAY=1
socket = r:TCP_NODELAY=1
verify = 2
CAfile = /etc/stunnel/stunnel.pem
client = yes
pid = /var/log/stunnel.pid
output = /var/log/stunnel.log
compression = zlib
ciphers = AES256-SHA
delay = no
failover = prio
sslVersion = TLSv1
fips = no

[proxy]
accept = 0.0.0.0:8088
connect = xxx.xxx.xxx.xxx:443
EOF
```
```
/usr/bin/stunnel
```
