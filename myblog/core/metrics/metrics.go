package metrics

import "github.com/gin-gonic/gin"

func InnerHandler(engine *gin.Engine) {
	innerGroup := engine.Group("/")
	innerGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
