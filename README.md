<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">🥳 ADMIN-SERVER-EXAMPLE</h1>
<h3 align="center">使用 Go 开发的后台管理系统后端示例</h3>

<p align="center">
  <a>
    <img src="https://img.shields.io/badge/-Golang 1.20-blue?style=flat-square&logo=go&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-Gin 1.9.1-blue?style=flat-square&logo=gin&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-MySQL-blue?style=flat-square&logo=mysql&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-Redis-c14438?style=flat-square&logo=redis&logoColor=white&link=mailto:ezops.cn@gmail.com" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-Minio-green?style=flat-square&logo=minio&logoColor=white&link=mailto:ezops.cn@gmail.com" alt="">
  </a>
</p>

<hr>

### 🤔 技术栈

- [x] Go：Google 开发的开源编程语言，诞生于 2006 年 1 月 2 日 15 点 4 分 5 秒 [:octocat:](https://github.com/golang/go)
- [x] Cobra：CLI 开发参数处理工具 [:octocat:](https://github.com/spf13/cobra)
- [x] Embed：go 1.16 新增的文件嵌入属性, 轻松将静态文件打包到编译后的二进制应用中
- [x] Gin：用 Go 编写的高性能 HTTP Web 框架 [:octocat:](https://github.com/gin-gonic/gin)
- [x] Viper：配置管理工具, 支持多种配置文件类型 [:octocat:](https://github.com/spf13/viper)
- [x] Zap：提供快速、结构化、分级的日志记录 [:octocat:](https://pkg.go.dev/go.uber.org/zap)
- [x] Lumberjack：日志滚动切割工具 [:octocat:](https://github.com/natefinch/lumberjack)
- [x] Gorm：数据库 ORM 管理框架, 可自行扩展多种数据库类型 [:octocat:](https://gorm.io/gorm)
- [x] Carbon：简单、语义化且对开发人员友好的 datetime 包 [:octocat:](https://github.com/golang-module/carbon)
- [x] Redis：Redis 客户端 [:octocat:](https://github.com/redis/go-redis)
- [x] Sonic：字节开源的高性能 JSON 库 [:octocat:](https://github.com/bytedance/sonic)
- [x] Jwt：用户认证, 登入登出一键搞定 [:octocat:](https://github.com/appleboy/gin-jwt)
- [x] OPT：OPT 双因子认证库 [:octocat:](https://github.com/pquerna/otp)
- [x] Copier：结构体转换库 [:octocat:](https://github.com/jinzhu/copier)
- [x] GoMail：发送邮件 [:octocat:](https://github.com/itrepablik/gomail)
- [x] Casbin：一个强大的、高效的开源访问控制框架 [:octocat:](https://casbin.org/zh/docs/overview)

<br>

### ⚡ 开发依赖

* 项目开发所需依赖的第三方包安装方法：

```bash
# 命令行工具
go get -u github.com/spf13/cobra

# Golang web 开发框架
go get -u github.com/gin-gonic/gin

# YAML 配置文件解析成结构体
go get -u github.com/spf13/viper

# 日志
go get -u go.uber.org/zap

# 日志切割
go get -u github.com/natefinch/lumberjack

# 数据库 GORM
go get -u gorm.io/gorm

# MySQL 连接驱动
go get -u gorm.io/driver/mysql

# 日期时间
go get -u github.com/golang-module/carbon/v2

# Redis 客户端
go get -u github.com/redis/go-redis/v9

# Sonic Json 库
go get github.com/bytedance/sonic

# JWT 认证
go get -u github.com/appleboy/gin-jwt/v2

# OPT 认证
go get -u github.com/pquerna/otp

# 结构体转换
go get -u github.com/jinzhu/copier

# 邮件发送
go get -u github.com/itrepablik/gomail

# RBAC
go get -u github.com/casbin/casbin/v2
go get -u github.com/casbin/gorm-adapter/v3
```