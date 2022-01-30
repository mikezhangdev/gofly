package sample

import (
	"github.com/gin-gonic/gin"
	"myblog/api/sample"
	"myblog/pkg/incontenxt"
)

func NewSampleRouter(engine *gin.Engine){
	group := engine.Group("/sample")
	group.Use(func(c *gin.Context) {
		c.JSON(200,"use")
		//c.Next()
	})
	group.GET("/check", func(ginCtx *gin.Context) {
		// 请求ID 注入业务上下文 TODO
		sample.Check(ginCtx,incontenxt.BuildInContext(ginCtx))
	})
}
