# 技术栈

## 编程语言
- Go 1.23.0+ (go.mod 指定 toolchain go1.23.7)

## 主要依赖库
- `github.com/go-sql-driver/mysql` - MySQL 数据库驱动
- `github.com/urfave/cli/v2` - 命令行接口框架
- `go.uber.org/zap` - 结构化日志库
- `golang.org/x/text` - 文本处理（编码转换）
- `gopkg.in/yaml.v2` - YAML 配置文件解析
- `github.com/mattn/go-colorable` - 控制台彩色输出

## 数据库
- MySQL (支持老版本密码认证)

## 构建工具
- Make (Linux)
- build.bat (Windows)
- 支持 UPX 压缩可执行文件