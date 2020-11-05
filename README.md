# p2pNG Core
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fp2pNG%2Fcore.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fp2pNG%2Fcore?ref=badge_shield)

## 开发环境
- Requirements: 
    - 安装[Golang](https://golang.org/dl/) Version 1.15
    - 配置Go Modules环境变量
        - 方案1: 配置GOPROXY，例如使用https://goproxy.cn/
        - 方案2: 配置https_proxy，指向梯子的本地http代理端口
    - 建议使用[GoLand](https://www.jetbrains.com/go/) IDE
    - 安装**godoc**,`go get -u golang.org/x/tools/cmd/godoc`
    - 安装**golint**,`go get -u golang.org/x/lint/golint`
- Release Build:
    - [core-builder](https://github.com/p2pNG/core-builder)
- Generate Docs
    - 在项目跟目录下运行`godoc`，浏览器打开
    [localhost:6060](http://localhost:6060/pkg/github.com/p2pNG/core/)

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

## 开发流程
1. 首次开始代码工作的准备
    - 对于团队成员可直接将[p2pNG/core](https://github.com/p2pNG/core) 克隆到本地，`git clone https://github.com/p2pNG/core`
    - 对于非团队成员，需要先进行[Fork](https://github.com/p2pNG/core/fork) 操作，并将其克隆到本地
2. 每次开发前，从最新的远程master分支创建本地分支，
    ```shell script
    git checkout master
    git pull
    git checkout -b [YOUR_BRANCH_NAME] master
    ```
3. 对于团队成员，应注意如下几点
    - Commit Message中、Pull Request的标题中必须包含Jira中的任务编号，例如`[P2PNG-27] Configure Github Actions CI`
    - Branch命名也建议包含任务编号的数字部分，例如：`27-configure-ci`
    - 提交代码前，需要使用`golint ./...`确认代码不存在问题；使用`go fmt ./...`进行代码格式化
3. 在`[YOUR_BRANCH_NAME]`branch上进行开发，完成开发后首次push如下
    `git push --set-upstream origin [YOUR_BRANCH_NAME]`；
    后续如有更多commit需要提交，直接`git push`到此分支即可
4. 按照Console中提示，到GitHub中创建Pull Request，等待CI构建和检查、审阅、合并


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fp2pNG%2Fcore.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fp2pNG%2Fcore?ref=badge_large)