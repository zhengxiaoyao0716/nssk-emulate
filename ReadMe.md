# nssk-emulate
## NSSK协议模拟

***
### 使用说明

***
### 功能规格
> - master: 作为中央公信服务器的一台PC或一个接口
> - user: 多台PC或同一PC多端口
1. Master端提供User端之间的接入与认证，提供app的下载
2. User端可以发起或接受创建安全加密连接的请求
3. 遵循NSSK协议建立安全连接，展示认证流程
4. User之间安全连接建立后可以相互以密文通讯
5. 通讯内容采用AES对称加密，解密后展示在前端

***
### 技术规格
1. 客户端以web形式展现，由h5前端页面+本地server后端组成，后端Server基于`Macaron`框架
2. NSSK协议中对称加密采用AES算法，16位密钥，基于Go官方标准库`"crypto/aes"`

***
### 附件
[NSSK协议认证流程与原理](./.request/NSSK.md)
