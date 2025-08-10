# 项目运行说明

## 先决条件

- 已安装 Go 1.20+（建议最新稳定版本）
- 已配置好数据库（PostgreSQL/MySQL，根据项目需要）
- 项目源码已下载到本地

---

## 一、初始化模块和安装依赖

进入项目根目录，执行：

```bash
# 初始化 go.mod（如果还没有）
go mod init your-module-name
# 安装项目依赖（根据代码 import 自动下载缺失依赖）
go mod tidy
# 手动安装缺少的驱动（如果用到了）
go get gorm.io/driver/postgres
```

## 二、配置数据库
1. 准备数据库，确保数据库服务已启动

2. 创建数据库（例如 PostgreSQL）：

```bash
psql -U your_username -h localhost
# 进入 psql 交互终端后执行：
CREATE DATABASE chatroom;
\q
```
3. 修改项目中数据库连接配置文件（如 config.json 或环境变量）

## 三、运行项目
```bash
# 直接运行所有 go 文件
go run db.go model.go server.go main.go

# 或者编译后启动
go build -o chatroom
./chatroom
```

## 四、常用命令
```bash
# 下载所有依赖
go mod tidy
# 查看当前依赖列表
go list -m all
# 清理未使用的依赖
go mod tidy
# 手动安装新包
go get package-name
# 运行单个 go 文件
go run filename.go
```

## 五、注意事项
如果网络受限，可设置代理环境变量，例如：

```bash
export GOPROXY=https://goproxy.cn,direct
```