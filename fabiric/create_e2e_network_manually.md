```
cd $GOPATH/github.com/hyperledger/fabric
mkdir aberic
cp example/e2e_cli/configtx.yaml ./aberic
cp example/e2e_cli/crypto-config.yaml ./aberic
./release/linux-amd64/bin/cryptogen --config=./aberic/crypto-config.yaml
mkdir aberic/crypto-config/idemix
cd aberic/crypto-config/idemix
./release/linux-amd64/bin/idemixgen ca-keygen
./release/linux-amd64/bin/idemixgen signerconfig -u OU1 -e OU1 -r 1
cd -
./release/linux-amd64/bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./release/linux-amd64/bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/mychannel1.tx -channelID mychannel

```
