package user_mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
}

type UserInfo struct {
	Id         int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	PhoneNum   string `json:"phone_num"`
	Password   string `json:"password"`
	CreateTime int64  `json:"create_time"`
	ModifyTime int64  `json:"modify_time"`
	DeleteTime int64  `json:"delete_time"`
	IsDel      int32  `json:"is_del"`
}


func (u *UserMysql) AddUserInfo(db *gorm.DB, model *UserInfo) error {
	t := time.Now().Unix()
	model.CreateTime = t
	model.ModifyTime = t
	err := db.Table("user_info").Create(model).Error
	return err
}

type GetUserListIn struct {
	Id    int64
	Pager struct {
		Page     int
		PageSize int
	}
	WithCount bool
	PhoneNum string
}

func (u *UserMysql) GetUserList(db *gorm.DB, in *GetUserListIn) ([]*UserInfo, int64, error) {
	var ret []*UserInfo
	var totalRows int64
	db = db.Table("user_info")
	if in.Id != 0 {
		db = db.Where("id = ?", in.Id)
	}
	if in.PhoneNum != "" {
		db = db.Where("phone_num = ?", in.PhoneNum)
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
