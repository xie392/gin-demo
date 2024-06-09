# Gin Demo

# 学习顺序

- 1、[gin-hello](./gin-hello): 最简单的 gin 项目
- 2、[gin-params](./gin-params): 处理请求参数
- 4、[gin-middleware](./gin-middleware): 中间件配置
- 5、[gin-template](./gin-template): 模板渲染
- 6、[gin-session](./gin-session): 会话管理
- 7、[gin-mysql](./gin-mysql): Mysql + ORM 集成
- 8、[gin-jwt](./gin-jwt): JWT 集成
- 9、[gin-swagger](./gin-swagger): Swagger 集成

# 错误解决

> go: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 142.251.42.241:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

tips: 如果出现以上错误，可以配置代理

```shell
# 1、开启 go module
go env -w GO111MODULE=on

# 2、设置 GOPROXY
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

# 如果是 macos 或 linux，可以直接设置环境变量：
# export GOPROXY=https://mirrors.aliyun.com/goproxy/

# 3、设置 GOSUMDB
go env -w GOSUMDB=off
#  export GOSUMDB=off // macOS 或 Linux
```

# 公认的工程目录（DDD 架构）

- cmd：存放项目的可执行文件,可能有多个 main 文件
- internal：存放项目的内部代码，不对外公开
- pkg：存放项目的公开代码，可以被其他项目依赖
- config/configs/etc：存放项目的配置文件
- scripts：存放项目的脚本文件
- docs：存放项目的文档
- third_party：存放项目的第三方依赖库
- bin: 编译的二进制文件
- build: 持续集成相关的配置文件
- deploy: 发布相关的配置文件
- test: 单元测试相关的配置文件
- api: 开发 API 接口
- init: 初始化项目的配置文件

# 创建项目

```shell
gin mod init [project name]
```

# 安装 Gin

```shell
go get -u github.com/gin-gonic/gin
```

# 运行项目

```shell
gin run main.go
```

# go.work

```shell
go 1.22.4

use (
    ./gin-hello
    ./gin-hello
    ./gin-middleware
    ./gin-params
    ./gin-template
    ./gin-session
    ./gin-orm
    ./gin-mysql
)

// 每次修改 go.mod 后，需要执行 go mod tidy 命令更新依赖
// go work sync
```
