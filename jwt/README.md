#jwt-go

**什么是JWT**
- 全称JSON WEB TOKEN
- 一种后台不做存储的前端身份验证的工具
- 分为三部分 Header Claims Signature
- 举例类比生活：(买彩票中奖后去兑奖，只人票，不认人)

**如何创建一个JWT**
- 通常使用NewWithClaims 因为我们可以通过匿名结构体来实现Claims接口，从而可以携带自己的参数
- 两种常用Claims的实现方式
   + type MyClaims struct{} 匿名函数实现接口
   + jwt.MapClaims{} map形式
- 创建一个Token
   + token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)//创建  
     s, err := t.SignedString(mySigningKey)//签发
- HS256(对称加密), 
  RS256(非对称加密，需要一个结构体指针 三个属性 {ECDSAKeyD,ECDSAKeyX,ECDSAKeyY},ES256)     



**如何解析一个JWT**
- token,err:=jwt.ParseWithClaims(token,func)
   + token：我们拿到的token字符串(ss)
   + 我们用哪个claims结构体发的，这里就传入哪个结构体
   + func: 一个特殊的回调函数，需要固定接收*Token类型指针，返回一个i 和一个err，此时的i就是我们的密钥
- token是一个jwtToken类型的数据，我们需要的就是其中的Claims
- 对Claims进行断言，然后进行取用即可
