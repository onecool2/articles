1、对称加密算法：DES， AES。 
  优点：加密速度快，即用密码加密，然后用同样的密码解密，
  缺点：但是如何进行密码在网络上的传输很麻烦，因为密码传输明文会被泄露，密文会涉及2次解密，一样需要密码。

2、非对称加密：RSA，ECC
  优点：一对公私钥，公钥传输给别人，私钥留给自己（通过私钥可以得出公钥），公钥用于加密，私钥用于解密，反之亦然
  缺点：加解密速度慢，但是没有密钥传输的问题，区块链体系里用ECC椭圆曲线算法实现。
 ECC大致的含义是，在一个椭圆曲线上取一点G，然后乘以一个整数k，得出一点K，即K=Gk，且K还是在椭圆曲线上的点，K为公钥，k为私钥，如果给出k和G求K比较容易，
 如果给出K和G求k就比较难了。

3、 哈希和散列函数（hash），典型算法SHA MD5
md5的值是128bit， 4位二进制组成的16进制数，所以最后是32个字节
sha256是256bit，最后结果是64个字节

4、 数字签名首先，发送方获取原文的摘要并用私钥进行加密，然后把原文和加密后的摘要发送给接受方，接收方用公钥进行解密，根据原文得到摘要，然后比较两者是否一至，如果一致则没有篡改，从而验证发送方身份

5、数字证书包含：私钥，公钥，第三方机构给你的签名，有效期。
