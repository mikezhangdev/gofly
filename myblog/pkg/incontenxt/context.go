package incontenxt

import (
	"context"
	"github.com/gin-gonic/gin"
	"myblog/pkg/util"
)

func BuildInContext(ginCtx *gin.Context)context.Context{
	// 获取外部传入的参数构建内部context
	requestId := ginCtx.GetHeader("request_id")
	if requestId == ""{
		requestId = util.UUid()
	}
	// 将公共参数设置进内部context
	inCtx := context.Background()
	inCtx = context.WithValue(inCtx,"trace.id",requestId)
	return inCtx
}
