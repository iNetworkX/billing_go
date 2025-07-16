# 建议的开发命令

## 构建命令 (Linux)
```bash
# 构建当前平台版本
make

# 清理构建产物
make clean

# 构建并打包32位版本
make x32

# 构建并打包64位版本
make x64

# 构建所有架构版本
make all
```

## 构建命令 (Windows)
```batch
# 双击运行
build.bat
```

## 代码质量检查
```bash
# 安装 golint（如果未安装）
go install golang.org/x/lint/golint@latest

# 运行 lint 检查
golint ./...

# 运行 vet 检查
go vet ./...

# 格式化代码
go fmt ./...
```

## 运行命令
```bash
# 前台运行
./billing

# 后台守护进程运行
./billing up -d

# 停止服务
./billing stop

# 查看版本信息
./billing version

# 查看在线用户
./billing show-users
```

## Git 常用命令
```bash
git status
git add .
git commit -m "commit message"
git push
```