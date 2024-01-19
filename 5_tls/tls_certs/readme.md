```shell
# 创建根CA私钥和自签名证书
openssl req -x509 -newkey rsa:4096 -nodes -keyout ca.key -out ca.crt -days 3650

# 创建服务器私钥
openssl genpkey -algorithm RSA -out server.key

# 创建服务器证书请求
openssl req -new -key server.key -out server.csr

# 使用CA为服务器证书签名
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650
```

