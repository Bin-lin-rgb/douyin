package dao

import (
	"douyin/model"
)

//	公共函数

func (mgr manager) IsExist(username string) (model.Userinfo, error) {
	var userinfo model.Userinfo
	result := mgr.db.Where("name=?", username).Find(&userinfo)

	//不返回密码
	userinfo.Password = ""

	return userinfo, result.Error

}

func (mgr manager) GetUserInfo(userId int64) (model.Userinfo, error) {
	var userinfo model.Userinfo

	result := mgr.db.Select("id", "name", "follow_count", "follower_count", "is_follow").Find(&userinfo, userId)
	return userinfo, result.Error
}
