# p2pNG Core
## 开发环境
- Requirements: 
    - [Golang](https://golang.org/dl/) Version 1.15
- Release Build:
    - [core-builder](https://github.com/p2pNG/core-builder)

## 项目架构
```
- cmd       项目入口
    - commands  入口指令的定义
        - run //todo: 通过flags加载插件
- internal  内部使用的部分功能的包装
    - logging   日志记录的包装 //todo: 分流写入log文件
    - utils
- modules   提供Services使用的基础功能
    - certificate   用于测试的证书生成
    - database      用于本地储存的key-value DB
    - file_storage  本地文件的相关操作
    - listener      Http服务监听
    - request       Http请求发送   
- services  Http服务（以插件形式加载）
    - status        基础节点信息交互
    - discovery     用户发现：通过mDNS和DoH
    - transfer      Seed、FileBlock交换
    - manage        限制本地Loopback使用的Core实例管理
    - trust         证书的签发、用户认证：通过ACME和OCSP
    - ... 下面为下一版需求
    - legacy_http   提供传统Http服务
    - token         Token签发、交换、验证
    - traversal     兼容无公网的IPv4用户： 通过UPNP,STUN,Proxy
```