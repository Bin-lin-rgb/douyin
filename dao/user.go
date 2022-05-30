package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
)

func (mgr manager) Register(userinfo model.Userinfo) error {
	result := mgr.db.Create(&userinfo)
	return result.Error

}

// UserToAuthor 一个人访问另一个人的主页查询是否关注
//func (mgr manager) UserToAuthor(userinfo model.Userinfo) error {
//	result := mgr.db.Model(model.Follow{}).Where("author_id = ", userinfo.Id)
//	return result.Error
//
//}
