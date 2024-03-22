# http

## GET 与 POST

GET 可以缓存，POST 不行

## HTTP/1.0 与 HTTP/1.1

HTTP/1.0 短连接，一次请求-响应后断开

HTTP/1.1 默认长连接

## HTTP/1.1

请求收到响应后，才能发送下一次请求，可能导致`队头阻塞`
