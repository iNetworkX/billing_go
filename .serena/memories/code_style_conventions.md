# 代码风格和约定

## 命名规范
- 包名：小写，单词直接连接（如 `bhandler`, `services`）
- 函数/方法：首字母大写的驼峰命名（如 `NewServer`, `LoadServerConfig`）
- 变量：驼峰命名（如 `serverConfig`, `clientData`）
- 常量：驼峰命名（如 `detachedProcess`）

## 代码注释
- 使用中文注释，解释函数功能和重要逻辑
- 导出函数必须有注释，格式：`// FunctionName 功能说明`
- 示例：`// ShowVersionInfo 展示版本信息`

## 错误处理
- 使用标准 Go 错误处理模式：`if err != nil { return nil, err }`
- 错误信息包装使用中文描述

## 项目结构约定
- `cmd/` - 命令行命令实现
- `services/` - 核心服务逻辑
- `models/` - 数据模型和数据库操作
- `common/` - 共享类型和接口
- `bhandler/` - 业务处理器（各种请求的处理逻辑）

## 编码规范
- 使用 UTF-8 编码
- 支持 GBK 编码转换（游戏客户端兼容）