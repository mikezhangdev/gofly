package sample

import (
	"context"
	"github.com/gin-gonic/gin"
	"myblog/api"
	"myblog/core"
)

func Check(ginCtx *gin.Context, inCtx context.Context) {
	// 打印日志
	core.Logger.InfoF(inCtx,"aaaa %v",123)
	api.CommonResp(ginCtx, &api.DataStruct{
		Data: struct {
			Title string
		}{
			Title: "标题",
		},
	})

}

//type Field = zapcore.Field

