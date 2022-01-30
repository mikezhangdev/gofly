





### 配置文件

```
https://github.com/spf13/viper
go get github.com/spf13/viper
```



### 日志

```
https://pkg.go.dev/go.uber.org/zap#section-readme
go get -u go.uber.org/zap
```



### Mysql

```
https://github.com/go-gorm/gorm
https://gorm.io/
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```



### Redis

```
https://github.com/gomodule/redigo
https://pkg.go.dev/github.com/gomodule/redigo/redis

go get github.com/gomodule/redigo/redis
```



### Webserver

```
https://github.com/gin-gonic/gin

go get -u github.com/gin-gonic/gin
```



### 外部组件引用

```
uuid
github.com/google/uuid
```



### 接口文档

```
https://github.com/go-swagger/go-swagger
go get -u github.com/swaggo/swag/cmd/swag
https://goswagger.io/
https://swagger.io/specification/

```





### 项目结构

```
├ ▏▏│ ─ 


├── main.go  // 项目入口文件
│
├── router   // 路由层
│
├── api      // 提供api层
│
├── service  // 提供service层
│
├── model    // 提供输入输出结构体层
│
├── core
│	 ├── app
│    │    ├── app.go
│    │    ├── gin.go
│    │
│    ├── setup
│    │    ├── config.go
│    │    ├── log.go
│    │    ├── mysql.go
│    │    ├── redis.go
│    │
│    ├── setting
│    │    ├── setting.go
│    │
│    ├── logging
│    │
│    │
│    ├── metrics
│    │
│    │
│    ├── core.go

```



### 输入输出结构

#### 输入

- header公参

  | Authorization | 登录态鉴权字符串 |      |
  | :------------ | ---------------- | ---- |
  |               |                  |      |
  |               |                  |      |
  |               |                  |      |

  

#### 输出

```
{
	"code":200, // 返回状态吗 200代表返回成功 -1 代表通用异常状态码 具体业务自定义状态码根据具体场景定义
	"msg":"返回提示",  // 可以给到用户侧的提示
	"inner_msg":"内部错误提示",  // 用来作为方便排查问题 不出现在用户展示侧
	"data":{}, // 返回结构
	"ts":1641964668 // 服务器当前时间
	"request_id": "请求UUID" // 请求唯一UUID 日志上会携带改请求ID 可以通过请求ID看链路日志
}
```



### 用户模块





