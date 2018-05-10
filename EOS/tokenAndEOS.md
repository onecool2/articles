发行token
cleos wallet create  
cleos import <eosio private key>  
cleos set contract eosio contracts/eosio.bios/  
cleos create account eosio eos.store <owner key> <active key>  
cleos set contract eos.store contracts/eosio.token/  
cleos push action eos.store create '["eos.store","10000000 EOS",0,0,0]' -p eos.store  
cleos push action eos.store issue '["eos.store","100000 EOS","issue"]' -p eos.store  
cleos get currency balance eos.store eos.store EOS  
  
cleos push action eos.store transfer '["eos.store","eosio","100 EOS",""]' -p eos.stor

cleos get currency balance eos.store eosio
发行EOS官方币
cleos create account eosio eos.offical <owner key> <active key>  

cleos set contract  eos.offical contracts/eosio.system/ 

cleos push action eos.offical  issue '["eos.offical", "1000000.0000 EOS",""]' -p eos.offical 


