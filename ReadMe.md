# nssk-emulate
## NSSK协议模拟

***
### 使用文档
``` bat
Usage of nssk-emulate:
  -host string
        指定主机IP或域名 (default "localhost")
  -master string
        指定公信服务主机通讯地址
         * 当公信服务地址与本机服务地址相同时
         * (master == http://address:port)
         * 应用将作为公信服务端启动
        (default "http://localhost:5000")
  -port int
        指定端口号（-1表示自动检索） (default -1)
  -s string
        信号 run | scan (default "run")

```

***
### 功能规格
> - master: 作为中央公信服务器的一台PC|Server或一个端口
> - user: 多台PC|Server或同一PC|Server多端口
> - S：NSSK认证协议中充当中间服务器的角色
> - A: NSSK认证协议中发起建立连接请求，接受连接验证的角色
> - B: NSSK认证协议中接受安全连接请求，发起连接验证的角色

1. Master端提供User端之间的接入与认证，提供app的下载
2. User端可以发起或接受创建安全加密连接的请求
3. 遵循NSSK协议建立安全连接，展示认证流程
4. User之间安全连接建立后可以相互以密文通讯
5. 通讯内容采用AES对称加密，解密后展示在前端
6. Master自身也可以作为User，User可以与自己建立连接并通讯
7. S、A、B三端可以在同一个节点上同时演示，也可以分开演示

***
### 技术规格
1. 客户端以web形式展现，由h5前端页面+本地server后端组成，后端Server基于`Macaron`框架
2. NSSK协议中对称加密采用AES算法，16位密钥，基于Go官方标准库`"crypto/aes"`

***
### 附件
[NSSK协议认证流程与原理](./.request/NSSK.md)
