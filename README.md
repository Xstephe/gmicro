# gmicro
自研微服务框架gmicro重构电商系统

#### 介绍

【重构中】

基于自研gmicro电商系统的重构版本。

http服务和rpc服务采用ioc代码分层。

http服务 基于gin封装。rpc服务 基于gRPC封装。（第三方接口，借鉴kortos进行服务解耦 可使用其他进行二次封装）

自定义log包 基于zap封装、errors包。全适配官方的包

借鉴kratos服务注册进行解耦。
默认使用consul进行服务的注册、发现、负载

http服务的认证方式默认是：JWT

现在用的是cobra、pflag、viper进行配置管理
启动时可使用--help进行查看

启动后台服务命令：go run ./cmd/admin/admin.go -c=configs/admin/admin.yaml
可在GoLand中配置启动时Program arguments为：-c=configs/admin/admin.yaml

#### 软件架构


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

使用命令行可查看

