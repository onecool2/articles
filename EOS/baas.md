for name in msig token disco bios ram ramfee names stake saving bpay vpay unregd  
do  
  cleos -u http://nodeos:80 create account eosio eosio.$name EOS7FNpU6P3yhT7bkf3fs7aNUxb3yNbMXPbn4nsYgh3ZkEhtchEAH  
done  

for name in msig token  
do  
  cleos -u http://nodeos:80 set contract eosio.$name /eos_code/mainnet/eos/build/contracts/eosio.$name  
done  


cleos wallet import 5Kgh744izqDCNyuRZcBu4nMjgJbBwWBdqQZuQi53qPB9L69Cj9X

cleos -u http://nodeos:80 create account eosio eosio.token EOS6CYWY1q4AAsVV3rzcT1GGQmZcg7yDnA6jZx2KUjHvmZWPv8DQg EOS6CYWY1q4AAsVV3rzcT1GGQmZcg7yDnA6jZx2KUjHvmZWPv8DQg

cleos -u http://nodeos:80  set contract eosio.token  eos/build/contracts/eosio.token/

cleos -u http://nodeos:80  push action eosio.token create '["eosio","1000000000.0000 SYS",0,0,0]' -p eosio.token

cleos -u http://nodeos:80 push action eosio.token issue '["eosio","1000000000.0000 SYS","issue"]' -p eosio

cleos -u http://nodeos:80 get currency balance eosio.token eosio

cleos -u http://nodeos:80  create account eosio alex EOS6CYWY1q4AAsVV3rzcT1GGQmZcg7yDnA6jZx2KUjHvmZWPv8DQg EOS6CYWY1q4AAsVV3rzcT1GGQmZcg7yDnA6jZx2KUjHvmZWPv8DQg

cleos -u http://nodeos:80  push action eosio.token transfer '["eosio", "alex", "10000.0000 SYS", ""]' -p eosio

#####################################################################################
cleos -u http://nodeos:80 set contract eosio eos/build/contracts/eosio.system

cleos -u http://nodeos:80 system newaccount eosio bp1 EOS7FNpU6P3yhT7bkf3fs7aNUxb3yNbMXPbn4nsYgh3ZkEhtchEAH EOS7FNpU6P3yhT7bkf3fs7aNUxb3yNbMXPbn4nsYgh3ZkEhtchEAH --stake-net '500.00 SYS' --stake-cpu '500.00 SYS'

cleos -u http://nodeos:80 wallet import 5Je4PsFWinBPXgWQasz7usbv8MLu5AbkcjU41JQr4Suhh34yKPu

cleos -u http://nodeos:80 system regproducer bp1 EOS7FNpU6P3yhT7bkf3fs7aNUxb3yNbMXPbn4nsYgh3ZkEhtchEAH http://bp1:8888


cleos -u http://nodeos:80 system delegatebw eosio bp1 '100000000.0000 SYS' '50000000.0000 SYS' --transfer

cleos -u http://nodeos:80 system voteproducer prods bp1 bp1

cleos -u http://nodeos:80 wallet unlock -n default --password PW5KcZAxfwhycUyWAjxJzuKSAZfwUoRgj9soJgQ1H5AbQAmLk23Ew

cleos -u http://nodeos:80 push action eosio.token create '["eosio","1000000000.0000 SYS",0,0,0]' -p eosio.token
cleos -u http://nodeos:80 push action eosio.token issue '["eosio","1000000000.0000 SYS","issue"]' -p eosio

cleos -u http://nodeos:80 set contract eosio eos/build/contracts/eosio.system


# Load all keys from /root/baas/Key_file.txt
cleos -u http://nodeos:80 wallet import 5Je4PsFWinBPXgWQasz7usbv8MLu5AbkcjU41JQr4Suhh34yKPu

cleos -u http://nodeos:80 system newaccount eosio bp1 EOS6ZDFJbYMNKaee2PDvN7wAB8DCJYd7XKssMQhHBsc8mYQiTrXrC EOS6ZDFJbYMNKaee2PDvN7wAB8DCJYd7XKssMQhHBsc8mYQiTrXrC --stake-net '500.00 SYS' --stake-cpu '500.00 SYS' --buy-ram '100 SYS'
cleos -u http://nodeos:80 system newaccount eosio bp2 EOS8L76iViM7KFdNht8mY9cZMtwcxmRXNhVn6WFmkj8uJJSioYWvV EOS8L76iViM7KFdNht8mY9cZMtwcxmRXNhVn6WFmkj8uJJSioYWvV --stake-net '500.00 SYS' --stake-cpu '500.00 SYS' --buy-ram '100 SYS'
cleos -u http://nodeos:80 system newaccount eosio bp3 EOS8LA9qmotA2pve9HdW7uDyLdL22dwcch9Bvb126iNNfLuM3peGi EOS8LA9qmotA2pve9HdW7uDyLdL22dwcch9Bvb126iNNfLuM3peGi --stake-net '500.00 SYS' --stake-cpu '500.00 SYS' --buy-ram '100 SYS'
# Registry all bp user from key_file.txt
cleos -u http://nodeos:80 system regproducer bp1 EOS7FNpU6P3yhT7bkf3fs7aNUxb3yNbMXPbn4nsYgh3ZkEhtchEAH http://bp1:8888

cleos -u http://nodeos:80 system delegatebw eosio bp1 '100000000.0000 SYS' '50000000.0000 SYS' --transfer

cleos -u http://nodeos:80 system voteproducer prods bp1 bp1
