# hat_logic 游戏逻辑服务器

## 项目简介
`hat_logic` 是一个基于 Go 语言开发的游戏逻辑服务器，采用微服务架构，通过 Consul 进行服务发现和配置管理，使用 Protobuf 和 gRPC 实现高效通信，旨在为游戏业务提供稳定、高效的逻辑处理服务。

## 目录
- [项目简介](#项目简介)
- [功能模块](#功能模块)
- [架构设计](#架构设计)
- [本地启动](#本地启动)
- [开发工具](#开发工具)
- [部署](#部署)
- [版权声明](#版权声明)
- [联系我们](#联系我们)

## 功能模块
- **游戏逻辑模块**：处理游戏的核心业务逻辑，如玩家操作、游戏规则等。
- **通信模块**：通过 gRPC 提供对外的 RPC 接口，实现与其他服务的高效通信。
- **配置管理模块**：利用 Consul 动态管理配置信息，支持不重启服务更新配置。
- **数据存储模块**：负责存储游戏数据，如玩家信息、游戏状态等。
- **日志模块**：记录服务运行日志，便于问题排查和系统维护。

## 架构设计
- **微服务架构**：将游戏逻辑作为独立服务运行，通过 Consul 实现服务发现和配置管理，提高系统的可扩展性和可用性。
- **高效通信**：采用 Protobuf 和 gRPC，确保通信的高效性和低延迟。
- **动态配置管理**：借助 Consul 的 Key/Value 存储动态更新配置，减少服务重启。
- **代码生成模板**：使用 Goland 的 Live Templates 功能，快速生成模板代码，提高开发效率。

## 本地启动
1. **配置参数**：根据需要设置本地启动参数，如集群名称、服务名、节点 ID 等。示例：`hat_dev hat_logic 5 debug`。
2. **生成 Protobuf 文件**：
    - 存储 pb 生成命令：`protoc --go_out=../hat_logic/ proto/comm.proto && protoc --go_out=../hat_logic/ proto/player.proto`。
    - 请求 pb 生成：`protoc --go_out=plugins=kite:../hat_logic comm.proto && protoc --go_out=plugins=kite:../hat_logic logic.proto`。
    - 对外 rpc 接口生成：`protoc --go_out=plugins=kite:../hat_logic/ proto/api.proto`。
3. **启动服务**：运行相关命令启动游戏逻辑服务器。

## 开发工具
- **代码编辑器**：推荐使用 Goland，方便进行 Go 语言开发。
- **代码生成模板**：在 Goland 的 setting 中搜索 live temp，设置 apiMod 快速生产模版代码。

## 部署
- **线上部署**：将服务部署到服务器，配置好 Consul 和相关依赖。
- **错误日志查询路径**：`/home/dhcd/back/service_deploy/re_dev-re_logic-1/log`。

## 版权声明
本项目遵循 [MIT 许可证](LICENSE)，详情请参阅 [LICENSE](LICENSE) 文件。

## 联系我们
- **作者**：EddieChan1993
- **项目链接**：[hat_logic](https://github.com/EddieChan1993/hat_logic)
- **邮箱**：（如有需要，可添加邮箱地址）