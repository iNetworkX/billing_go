# 当前项目状态

## 版本信息
- 基于 Go 1.23.0+
- 主要功能：billing 验证服务器，用于游戏账号认证和计费管理

## 活跃开发状态
- 项目正在积极开发中
- 最近实现了 IP 信息监控功能
- 代码结构稳定，核心功能完整

## 主要组件
- **服务器核心** (`services/billing/`) - 完整的服务器实现
- **请求处理** (`bhandler/`) - 各种业务请求处理器
- **数据模型** (`models/`) - MySQL 数据库操作
- **命令行界面** (`cmd/`) - CLI 命令实现
- **共享类型** (`common/`) - 数据结构定义

## 运行支持
- 支持前台/后台运行模式
- 支持 systemd 服务
- 配置文件支持 YAML/JSON 格式
- 完整的日志系统

## 当前可用命令
- `billing up` - 启动服务器
- `billing stop` - 停止服务器
- `billing show` - 显示 IP 信息
- `billing show-users` - 显示用户状态
- `billing version` - 版本信息