# [https](https://www.ruanyifeng.com/blog/2014/02/ssl_tls.html)

## HTTP 与 HTTPS

HTTP 是明文传输

HTTPS 使用TLS加密传输，大多使用ECDHE密钥协商

## RSA密钥协商

### 握手

客户端发出请求 - ClientHello

服务端回应 - ServerHello，回应中包含服务器证书

客户端回应

服务端回应

### 如何加密

前两次通信，生成两个随机数，明文传输

第三次通信，生成一个随机数，使用证书中的公钥加密，服务器使用私钥解密 - `非对称加密`

使用三个随机数，以及约定的加密方法，生成`会话密钥`

握手之后，使用`会话密钥`加密HTTP数据包 - `对称加密`
