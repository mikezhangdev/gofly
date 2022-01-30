package router

import (
	"github.com/gin-gonic/gin"
	swagger_router "myblog/api/swagger"
	"myblog/router/sample"
	"myblog/router/user_router"
)

func GinHandlerRouter(engine *gin.Engine)error{
	sample.NewSampleRouter(engine)
	user_router.NewUserRouter(engine)
	swagger_router.RegisterSwagger(engine)
	return nil
}
