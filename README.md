# gmicro
自研微服务框架 gmicro 重构电商系统

## 项目介绍

gmicro 是一个基于 Go 语言的微服务框架，旨在重构电商系统，提供高性能、可扩展的服务架构。

## 技术栈

- **框架与库**：
  - cobra：命令行应用程序管理
  - grpc：高性能跨语言的远程过程调用 (RPC) 框架
  - protoc 插件开发：用于生成 gRPC 代码
  - ast 开发：抽象语法树操作
  
- **设计与架构**：
  - errors 设计：自定义错误处理
  - log 日志设计：基于 zap 的日志框架
  - 三层代码架构：分层设计提升可维护性

- **分布式系统**：
  - dtm：分布式事务管理框架
  - kratos 与 go-zero：微服务框架
  - open-telemetry：分布式追踪与监控
  - prometheus & grafana：监控与可视化工具

- **测试与模拟**：
  - sqlmock、gomock、gomonkey：用于单元测试与模拟
  - 模糊测试：提高代码的健壮性

- **DevOps**：
  - CI/CD：持续集成与持续交付
  - k8s & kubesphere：容器编排与管理
  - harbor：云原生的容器镜像仓库

## 软件架构

- **服务架构**：
  - HTTP 服务与 RPC 服务采用 IOC 代码分层设计。
  - HTTP 服务基于 Gin 封装，支持中间件与路由配置。
  - RPC 服务基于 gRPC 封装，支持多种协议。

- **服务解耦**：
  - 借鉴 Kortos 进行服务解耦，支持多种第三方接口的整合。

- **服务注册与发现**：
  - 默认使用 Consul 进行服务的注册、发现与负载均衡。

- **认证与安全**：
  - HTTP 服务的认证方式默认是 JWT（JSON Web Token）。
