package dao

import (
	"douyin/model"
	"strings"
)

//	公共函数

func (mgr manager) IsExist(username string) (model.Userinfo, error) {
	var userinfo model.Userinfo
	result := mgr.db.Model(&model.Userinfo{}).Where("name=?", username).Find(&userinfo)

	return userinfo, result.Error
}

func (mgr manager) GetUserInfo(userId int64) (model.Userinfo, error) {
	var userinfo model.Userinfo

	result := mgr.db.Select("id", "name", "follow_count", "follower_count", "is_follow").Find(&userinfo, userId)
	return userinfo, result.Error
}

func StrBuilder(first string, finalName string) string {
	var builder strings.Builder

	builder.WriteString(first)
	builder.WriteString(finalName)

	return builder.String()
}
