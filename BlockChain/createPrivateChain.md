#### 1. Download geth binary which maches your platform from https://ethereum.github.io/go-ethereum/downloads/
   ```
   mkdir geth 
   wget https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-arm64-1.8.2-b8b9f7f4.tar.gz
   tar xvf geth-alltools-linux-arm64-1.8.2-b8b9f7f4.tar.gz -C /usr/bin
   ```
#### 2. Create a directory for your ethereum data and create account on each node
   ```
   mkdir testnet && cd testnet
   geth --datadir node account new
   ```
#### 3. Generate genesis.json
   ```
   puppeth
   ```
   
#### 4. Copy genesis.json to each node and run :
   ```
   geth --datadir node init genesis.json
   echo "your account password" > node/password
   geth --datadir node0 --port 30000 --nodiscover --unlock '0' --password ./node/password console
   ```
   
#### 5. Through consol found "enode" on each node and put them into static-nodes.json
    ```
    > admin.nodeInfo.enode
    ```
    ```               "enode://92e275a3f00d3688923dbf43105f86c4ec9c74e22efea114d08cfbe45eddf6ec25d6ca6b9e1251445127cbf31d4c9096739bf64a47624f0ef0fb30d87609f432@[::]:30000?discport=0"
    cat static-nodes.json 
    [
      "enode://7e97ed5dfdf1d01652e4fcdb27d2c598b30d228de5369b61fffeeec2a1956b176e5af5298543bbf45a4852a669cd7f4feeb031fdf1bef69de0817642e15a24b8@172.xx.xx.xx:30000?discport=0",
      "enode://92e275a3f00d3688923dbf43105f86c4ec9c74e22efea114d08cfbe45eddf6ec25d6ca6b9e1251445127cbf31d4c9096739bf64a47624f0ef0fb30d87609f432@172.xx.xx.xx:30000?discport=0"
    ]
    ```
#### 6. Copy static-nodes.json to "node" directory on each node and relanch geth 
   ```
   copy static-nodes node/
   geth --datadir node0 --port 30000 --nodiscover --unlock '0' --password ./node/password console
   run "admin.peers" to verify the operation whether is success.
   ```
### 7. mine.start()   
