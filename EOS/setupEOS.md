分章分为3个部分
1）介绍一下区块链和EOS的大致原理以及BP节点的角色职责

  关于代码的获得和编译过程，可以参照官方wiki的介绍，不作为本文重点，下面只给出主要执行步骤：
  EOS代码获得和编译过程：
  代码的获取可以从git clone https://github.com/EOSIO/eos --recursive
  cd eos
  ./eosio_build.sh
  
2）技术选型，以及大致部署步骤
  0.docker run的方式
  1.多节点共用一套公私钥，大致配置方式结果及解决办法
   
  2.两节点通过keep a live
  3.通过第三节点监控
  4.通过openshift部署

3）各种配置方案优缺点及总结
