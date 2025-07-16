# 项目结构

## 顶层目录结构
```
billing_go/
├── cmd/                    # 命令行命令实现
│   ├── app_command.go     # 主命令入口
│   ├── up.go              # 启动命令
│   ├── stop.go            # 停止命令
│   ├── version.go         # 版本信息命令
│   └── show_users.go      # 显示在线用户命令
├── services/              # 核心服务逻辑
│   ├── billing/           # billing 服务器实现
│   ├── handle/            # 连接处理
│   └── *.go               # 各种服务工具函数
├── models/                # 数据模型
│   ├── account.go         # 账号模型
│   ├── check_login.go     # 登录验证
│   ├── register_account.go # 账号注册
│   └── *.go               # 其他数据库操作
├── common/                # 共享类型定义
│   ├── billing_packet.go  # 数据包定义
│   ├── server_config.go   # 服务器配置
│   ├── client_info.go     # 客户端信息
│   └── *.go               # 其他共享类型
├── bhandler/              # 业务处理器
│   ├── login_handler.go   # 登录处理
│   ├── register_handler.go # 注册处理
│   ├── query_point_handler.go # 查询点数
│   └── *.go               # 其他业务处理器
├── main.go                # 程序入口
├── go.mod                 # Go 模块定义
├── go.sum                 # 依赖版本锁定
├── Makefile               # Linux 构建脚本
├── build.bat              # Windows 构建脚本
├── config.yaml            # 配置文件示例
├── billing.service        # systemd 服务配置
└── README.md              # 项目说明文档
```

## 核心组件说明
- **cmd**: CLI 命令实现，使用 urfave/cli 框架
- **services**: 核心服务逻辑，包括服务器启动、配置加载、数据包处理等
- **models**: 数据库模型和操作，主要处理账号相关数据
- **common**: 共享的数据结构和接口定义
- **bhandler**: 各种业务请求的处理逻辑，每个 handler 处理特定的请求类型