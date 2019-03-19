1. wget https://github.com/ipfs/go-ipfs/releases/download/v0.4.19/go-ipfs_v0.4.19_linux-amd64.tar.gz

2. mkdir ~/ipfs; mv go-ipfs_v0.4.19_linux-amd64.tar.gz ~/ipfs

3. cd ~/ipfs/; tar xvf go-ipfs_v0.4.19_linux-amd64.tar.gz

4. ./install.sh 

5.  ipfs init;ipfs bootstrap rm --all //初始化和移除默认的bootstrap节点

6. 两个节点启动ipfs daemon，此时因该可以找到对方，
还需注意一点，如果报告ipfs daemon启动失败，报告找不到
可以添加如下到~/.ipfs/config 
	"Access-Control-Allow-Methods": [
				"PUT",
				"GET",
				"POST"
			],
			"Access-Control-Allow-Origin": [
				"http://192.168.2.14:5001",
				"http://127.0.0.1:5001",
				"https://webui.ipfs.io"
			]
或者执行
 ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["http://192.168.2.14:5001", "http://127.0.0.1:5001", "https://webui.ipfs.io"]'
 ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "GET", "POST"]'
