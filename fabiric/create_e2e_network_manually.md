```
cd $GOPATH/github.com/hyperledger/fabric
mkdir aberic
cp example/e2e_cli/configtx.yaml ./aberic
cp example/e2e_cli/crypto-config.yaml ./aberic
./release/linux-amd64/bin/cryptogen --config=./aberic/crypto-config.yaml

```
