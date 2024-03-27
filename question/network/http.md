# http

## GET 与 POST

GET 可以缓存，POST 不行

## HTTP/1.0 与 HTTP/1.1

HTTP/1.0 短连接，一次请求-响应后断开

HTTP/1.1 默认长连接

## HTTP/1.1 与 HTTP/2

HTTP/1.1 `http队头阻塞`，必须按照请求的顺序返回响应

HTTP/2

- 对请求头做了压缩，不重复发送相同字段
- 通过 Stream ID 标识消息，请求响应可以乱序 - 解决`http队头阻塞`
- 服务器主动推送资源

## HTTP/2 与 HTTP/3

HTTP/2 `TCP队头阻塞`
