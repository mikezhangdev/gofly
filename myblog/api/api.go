package api

import (
	"github.com/gin-gonic/gin"
	"myblog/pkg/util"
	"time"
)

//{
//"code":200, // 返回状态吗 200代表返回成功 -1 代表通用异常状态码 具体业务自定义状态码根据具体场景定义
//"msg":"返回提示",  // 可以给到用户侧的提示
//"inner_msg":"内部错误提示",  // 用来作为方便排查问题 不出现在用户展示侧
//"data":{}, // 返回结构
//"ts":1641964668 // 服务器当前时间
//"request_id": "请求UUID" // 请求唯一UUID 日志上会携带改请求ID 可以通过请求ID看链路日志
//}

type DataStruct struct{
	Code int64 `json:"code"`
	Msg string `json:"msg"`
	InnerMsg string `json:"inner_msg"`
	Data interface{} `json:"data"`
	Ts int64 `json:"ts"`
	RequestId string `json:"request_id"`
}

func CommonResp(ginCtx *gin.Context,d *DataStruct){
	if d.Code == 0{
		d.Code = 200
	}
	if d.Msg == ""{
		if d.Code == 200{
			d.Msg = "请求成功"
		}else{
			d.Msg = "请求失败"
		}
	}
	if d.RequestId == ""{ // 请求ID应由上层context注入 TODO
		d.RequestId = util.UUid()
	}
	d.Ts = time.Now().UnixNano()
	//ginCtx.Abort()
	ginCtx.JSON(200,d)
	ginCtx.Next()
	//ginCtx.AbortWithStatus(200)
}
