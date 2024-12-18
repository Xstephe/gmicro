# gmicro
自研微服务框架 gmicro 重构电商系统

## 项目介绍

**gmicro** 是一个基于 Go 语言的微服务框架，旨在重构电商系统，提供高性能、可扩展的服务架构。

## 技术栈

- **框架与库**：
  - **cobra**：命令行应用程序管理。
  - **gRPC**：高性能跨语言的远程过程调用 (RPC) 框架。
  - **protoc 插件开发**：用于生成 gRPC 代码。
  - **AST 开发**：支持抽象语法树的操作。
  - **自定义错误处理**：重新封装了 errors 包。
  - **日志框架**：基于 zap 的日志框架重新封装。

- **微服务框架**：
  - **kratos & go-zero**：用于自研 gmicro 微服务框架的基础。

- **监控与追踪**：
  - **OpenTelemetry**：集成进行分布式追踪与监控。
  - **Prometheus & Grafana**：用于系统监控与数据可视化。

- **DevOps**：
  - **CI/CD**：支持持续集成与持续交付。
  - **Kubernetes & KubeSphere**：容器编排与管理。
  - **Harbor**：云原生的容器镜像仓库。

- **设计模式**：
  - 采用工厂模式，提升代码的可维护性和扩展性。

## 架构设计

- **服务架构**：
  - HTTP 服务与 RPC 服务采用 IOC（控制反转）代码分层设计。
  - HTTP 服务基于 Gin 封装，支持中间件与路由配置。
  - RPC 服务基于 gRPC 封装，兼容多种协议。

- **服务解耦**：
  - 借鉴 Kortos 实现服务解耦，支持多种第三方接口的整合。

- **服务注册与发现**：
  - 默认使用 Consul 进行服务的注册、发现与负载均衡，确保服务的高可用性。

- **认证与安全**：
  - HTTP 服务默认采用 JWT（JSON Web Token）进行用户认证，保障系统安全性。

