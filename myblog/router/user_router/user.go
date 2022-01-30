package user_router

import (
	"github.com/gin-gonic/gin"
	user_api "myblog/api/user"
	"myblog/pkg/incontenxt"
)

func NewUserRouter(engine *gin.Engine){
	group := engine.Group("/user")
	group.POST("/send_sms", func(ginCtx *gin.Context) {
		// 请求ID 注入业务上下文 TODO
		user_api.SendSms(ginCtx,incontenxt.BuildInContext(ginCtx))
	})
	group.POST("/register", func(ginCtx *gin.Context) {
		// 请求ID 注入业务上下文 TODO
		user_api.Register(ginCtx,incontenxt.BuildInContext(ginCtx))
	})
}

