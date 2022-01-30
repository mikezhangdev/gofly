package user_api

import (
	"context"
	"fmt"
	"myblog/api"

	//"encoding/json"
	"github.com/gin-gonic/gin"
	"myblog/core"
	user_service "myblog/service/user"
)

//phoneNUm string, deviceId string

type SendSmsIn struct {
	PhoneNum string `json:"phone_num"`
	DeviceId string `json:"device_id"`
}

func SendSms(ginCtx *gin.Context, inCtx context.Context) {
	// 参数验证
	var req SendSmsIn
	err := ginCtx.BindJSON(&req)
	if err != nil {
		core.Logger.InfoF(inCtx, "SendSms er %v", err)
		api.CommonResp(ginCtx, &api.DataStruct{
			Code:     -1,
			Msg:      "请求失败",
			InnerMsg: fmt.Sprintf("%v", err),
		})
		return
	}
	if req.DeviceId == "" || req.PhoneNum == "" {
		api.CommonResp(ginCtx, &api.DataStruct{
			Code:     -1,
			Msg:      "请求失败",
			InnerMsg: "参数不全",
		})
		return
	}
	userSrv := user_service.NewUser()
	err = userSrv.SendSms(inCtx, req.PhoneNum, req.DeviceId)
	if err != nil {
		api.CommonResp(ginCtx, &api.DataStruct{
			Code:     -1,
			Msg:      "请求失败",
			InnerMsg: fmt.Sprintf("%v", err),
		})
		return
	}
	api.CommonResp(ginCtx, &api.DataStruct{
	})
	return

}

type RegisterReq struct {
	PhoneNum string `json:"phone_num"`
	DeviceId string `json:"device_id"`
	RandCode string `json:"rand_code"`
	Password string `json:"password"`
}


// @Tags    用户模块
// @Summary 用户注册
// @Description 用户注册
// @Accept  json
// @Produce  json
// @param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param  body body user_api.RegisterReq true "交款查询参数"
// @Success 200 {object} api.DataStruct{data=user_service.RegisterOut}
// @Failure 500 {object} api.DataStruct
// @Router /user/register [post]
func Register(ginCtx *gin.Context, inCtx context.Context) {
	// 参数验证
	var req RegisterReq
	err := ginCtx.BindJSON(&req)
	if err != nil {
		api.CommonResp(ginCtx, &api.DataStruct{
			Code:     -1,
			Msg:      "注册失败",
			InnerMsg: fmt.Sprintf("invalid params"),
		})
		return
	}

	userService := user_service.NewUser()
	ret, err := userService.Register(inCtx, req.PhoneNum, req.Password, req.DeviceId, req.RandCode)
	if err != nil {
		api.CommonResp(ginCtx, &api.DataStruct{
			Code:     -1,
			Msg:      "注册失败",
			InnerMsg: fmt.Sprintf("%v", err),
		})
		return
	}
	api.CommonResp(ginCtx, &api.DataStruct{
		Data: ret,
	})
	return
}
