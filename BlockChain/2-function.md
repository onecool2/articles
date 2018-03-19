pragma solidity ^0.4.4;

contract DemoTypes {
  /*uint public resultOfF;*/
  //event LogUint(uint value);
  //event LogInt (int value);


  function f(uint a) public returns (uint b)
  {
    uint result = a * 8;
    var resultOfF = result; //debug used only;
    //LogUint(result);
    return result;
  }

  /*输入长度宽度，返回面积*/
  function f2(int width, int height) returns (int square) {
    if (width < 0 || height < 0) throw;
    int result = width * height;
    //LogInt(result);
    return result;
  }

  /*输入N，计算N的阶乘，循环实现*/
  function f3(uint n) returns (uint jiecheng) {
    if (n == 0) throw; uint result = 1;
    for (uint i=1; i<=n; i++) {
      result *= i;
    }
    return result;
  }

  /*计算从1到N的求和*/
  function f4(uint n) returns (uint sum) {
    if (n == 0) throw; uint result = 0;
    for (uint i=0; i<=n; i++) {
      result +=i;
    }
    return result;
  }
}
