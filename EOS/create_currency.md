cleos wallet create

cleos set contract eosio ../../contracts/eosio.bios -p eosio

./cleos create key  # OwnerKey

./cleos create key  # ActiveKey

./cleos wallet import <private-OwnerKey>

./cleos wallet import <private-ActiveKey>

cleos create account eosio currency <public-OwnerKey> <public-ActiveKey>

cleos get account currency（可选）

cleos get code currency（可选）

cleos set contract currency ../../contracts/currency
第一个currency是部署这个合约的账户名，第二个是合约名，这两个可以不一样

cleos push action currency create '{"issuer":"currency","maximum_supply":"1000000.0000 MGD","can_freeze":"0","can_recall":"0","can_whitelist":"0"}' --permission currency@active
MGD是代币名称，这一步是创建
cleos push action currency issue '{"to":"currency","quantity":"1000.0000 MGD","memo":""}' --permission currency@active
这一步是发行（issue）

cleos get table currency currency accounts
查看账户余额

cleos push action currency transfer '{"from":"currency","to":"eosio","quantity":"20.0000 MGD","memo":"my first transfer"}' --permission currency@active
转账

