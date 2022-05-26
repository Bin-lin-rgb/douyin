package dao

import (
	"douyin/model"
)

type UserInfoResponse struct {
	Id            int64  `json:"id,omitempty"  gorm:"primary_key;column:id;type:int;not null" `
	Name          string `json:"name,omitempty" gorm:"column:name;type:varchar;not null" `
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

//	公共函数

func (mgr manager) IsExist(username string) (model.Userinfo, error) {
	var userinfo model.Userinfo
	result := mgr.db.Where("name=?", username).Find(&userinfo)
	//result := mgr.db.Model(&userinfo).Where("name=?", username).Find(&UserInfoResponse{})
	//result := mgr.db.Select("name", "follow_count", "follower_count", "is_follow").Find(&userinfo)
	return userinfo, result.Error

}

func (mgr manager) GetUserInfo(userId int64) (model.Userinfo, error) {
	var userinfo model.Userinfo
	//result := mgr.db.Model(&userinfo).Find(&UserInfoResponse{}, userId)
	result := mgr.db.Select("id", "name", "follow_count", "follower_count", "is_follow").Find(&userinfo, userId)
	return userinfo, result.Error
}
