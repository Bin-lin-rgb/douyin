package dao

import (
	"douyin/model"
)

func (mgr manager) Register(userinfo model.Userinfo) error {
	result := mgr.db.Create(&userinfo)
	return result.Error
}
