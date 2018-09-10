1 lsb_release -a  
2 download docker CE from "https://download.docker.com/linux/ubuntu/dists/xenial/pool/stable/amd64/"  
3 sudo dpkg -i ./docker-ce_*.deb  
4 sudo apt-get install docker-compose  
5 sudo apt-get install golang-1.10 
6 ln -s /usr/lib/go-1.10/bin/go /usr/local/bin/go
7 sudo apt-get install curl 
8 curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -  
9 sudo apt-get install nodejs  
10 sudo apt-get install -y build-essential  
11 nodejs --version  
12 sudo npm install npm -g  
13 sudo usermod -aG docker alex  
14 systemctl unmask docker   
14.5 set your https proxy and add the  $HOME/bin into your $PATH on bashrc  
15 curl -sSL https://goo.gl/kFFqh5 | bash -s 1.0.6

