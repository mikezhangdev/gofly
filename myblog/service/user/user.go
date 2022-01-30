package user_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"myblog/core"
	"myblog/pkg/const_key"
	"myblog/pkg/util"
	"myblog/repository/user_mysql"
	"time"
)

const MaxSendNum = 10

type User struct {
}

func NewUser() *User {
	return &User{}
}

type SmsInfo struct {
	SendNum  int64
	RandCode string
}

func (u *User) SendSms(ctx context.Context, phoneNUm string, deviceId string) error {
	smsCacheKey := const_key.UserSms + phoneNUm
	cacheCon := core.RedisPool.Get()
	smsCacheStr, err := redis.String(cacheCon.Do("GET", smsCacheKey))
	if err != nil && err != redis.ErrNil {
		return err
	}
	smsInfo := SmsInfo{}
	if smsCacheStr != "" {
		err = json.Unmarshal([]byte(smsCacheStr), &smsInfo)
	}
	if smsInfo.SendNum > MaxSendNum {
		return fmt.Errorf("sms send over max num 10")
	}
	// 检测短信是否合法 发送次数 设备ID
	// 生成短信验证码
	randCode := util.GenRandStr(0, 4)
	// 调用短信厂商发送验证码
	// 验证码记录
	// 刷新发送验证码统计
	smsInfo.SendNum = smsInfo.SendNum + 1
	smsInfo.RandCode = randCode
	smsInfoByte, err := json.Marshal(smsInfo)
	if err != nil {
		return err
	}
	smsCacheVal := string(smsInfoByte)
	_, err = redis.String(cacheCon.Do("SET", smsCacheKey, smsCacheVal, "EX", 360))
	if err != nil {
		return err
	}
	// 返回
	return nil
}

/**

CREATE TABLE IF NOT EXISTS `user_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone_num` varchar(36) NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(36) NOT NULL DEFAULT '' COMMENT '密码',
  `create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modify_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
  `delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除@：0否@：1是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户信息表';


CREATE TABLE IF NOT EXISTS `user_login` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `device_id` varchar(36) NOT NULL DEFAULT '' COMMENT '设备唯一码',
  `access_token` varchar(36) NOT NULL DEFAULT '' COMMENT 'access_token',
  `refresh_token` varchar(36) NOT NULL DEFAULT '' COMMENT 'refresh_token',
  `access_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后刷新access_token时间',
  `modify_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
  `delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除@：0否@：1是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户登录表';
*/

func (u *User) Login(ctx context.Context, phoneNUm, randCode string) {

}

type RegisterOut struct {
	AccessToken  string
	RefreshToken string
	ExpireTime   int64
	Status       int32 // 1为正常 2 为该账户已完成注册走 登录流程 -1 通用异常
}

func (u *User) Register(ctx context.Context, phoneNUm, passWord, deviceId string, randCode string) (*RegisterOut, error) {
	db := core.MysqlConn
	// 检测手机号是否已完成注册 如果已注册 返回提示 走登录流程
	userMysql := user_mysql.NewUserMysql()
	uinfoTmp, _, err := userMysql.GetUserList(db, &user_mysql.GetUserListIn{
		PhoneNum: phoneNUm,
	})
	if err != nil {
		return &RegisterOut{Status: 2}, nil
	}
	if len(uinfoTmp) > 0 {
		return nil, fmt.Errorf("acount has exsist")
	}
	// 验证码是否有效
	smsCacheKey := const_key.UserSms + phoneNUm
	cacheCon := core.RedisPool.Get()
	smsCacheStr, err := redis.String(cacheCon.Do("GET", smsCacheKey))
	if err != nil && err != redis.ErrNil {
		return nil, err
	}
	smsInfo := SmsInfo{}
	if smsCacheStr != "" {
		err = json.Unmarshal([]byte(smsCacheStr), &smsInfo)
	}
	if smsInfo.RandCode != randCode {
		return &RegisterOut{Status: -1}, fmt.Errorf("randCode not right")
	}
	// 添加用户信息
	md5PassWord,err := util.GenMd5Str(passWord)
	if err != nil{
		return &RegisterOut{Status: -1}, fmt.Errorf("randCode not right")
	}
	userInfoModel := &user_mysql.UserInfo{
		PhoneNum:   phoneNUm,
		Password:   md5PassWord,
	}
	err = userMysql.AddUserInfo(db, userInfoModel)
	if err != nil {
		return &RegisterOut{Status: -1}, fmt.Errorf("register err")
	}
	userId := userInfoModel.Id
	log.Printf("userId: %v", userId)
	// 刷新登录token信息
	accessToken := util.UUid()
	RefreshToken := util.UUid()
	err = userMysql.AddLoginInfo(db, &user_mysql.UserLogin{
		UserId:       userId,
		DeviceId:     deviceId,
		AccessToken:  accessToken,
		RefreshToken: RefreshToken,
		AccessTime:   time.Now().Unix(),
	})
	if err != nil {
		return nil, fmt.Errorf("register err")
	}
	return &RegisterOut{
		AccessToken:  accessToken,
		RefreshToken: RefreshToken,
		ExpireTime:   3600,
		Status:       1,
	}, nil
}

func (u *User) GetUserInfoByAccessToken() {

}
