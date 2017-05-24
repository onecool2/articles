# On EC2:
```
yum install squid openssl openssl-devel
yum install stunnel
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
```
# On your local server:
scp stunnel.pem to your local server
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
```
