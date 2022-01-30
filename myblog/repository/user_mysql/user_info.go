package user_mysql

import (
	"gorm.io/gorm"
	"time"
)

type UserLogin struct {
	Id           int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserId       int64  `json:"user_id"`
	DeviceId     string `json:"device_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessTime   int64  `json:"access_time"`
	CreateTime   int64  `json:"create_time"`
	ModifyTime   int64  `json:"modify_time"`
	DeleteTime   int64  `json:"delete_time"`
	IsDel        int32  `json:"is_del"`
}

func (u *UserMysql) AddLoginInfo(db *gorm.DB, model *UserLogin) error {
	where := map[string]interface{}{
		"user_id": model.UserId,
	}
	t := time.Now().Unix()
	model.CreateTime = t // 抽出来到公共出来时间字段 TODO
	model.ModifyTime = t
	err := db.Table("user_login").Where(where).Create(model).Error
	return err
}

type GetUserLoginListIn struct {
	Id     int64
	UserId int64
	Pager  struct {
		Page     int
		PageSize int
	}
	WithCount bool
}

func (u *UserMysql) GetUserLoginList(db *gorm.DB, in *GetUserLoginListIn) ([]*UserLogin, int64, error) {
	var ret []*UserLogin
	var totalRows int64
	db = db.Table("user_login")
	if in.Id != 0 {
		db = db.Where("id = ?", in.Id)
	}
	if in.UserId != 0 {
		db = db.Where("user_id = ?", in.UserId)
	}
	if in.WithCount {
		db = db.Count(&totalRows)
	}
	if in.Pager.Page <= 0 {
		in.Pager.Page = 1
	}
	if in.Pager.PageSize <= 0 {
		in.Pager.PageSize = 1
	}
	err := db.Offset((in.Pager.Page - 1) * in.Pager.PageSize).Limit(in.Pager.PageSize).Find(&ret).Error
	return ret, totalRows, err
}
