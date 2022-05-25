package dao

import (
	"douyin/model"
)

//	公共函数

func (mgr manager) IsExist(username string) (model.Userinfo, error) {
	var t_userinfo model.Userinfo
	result := mgr.db.Where("name=?", username).Find(&t_userinfo)
	return t_userinfo, result.Error
}

func (mgr manager) GetUserInfo(userId int64, userinfo *model.Userinfo) error {
	result := mgr.db.First(&userinfo, userId)
	return result.Error
}
