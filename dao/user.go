package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
)

func (mgr manager) Register(userinfo model.Userinfo) error {
	result := mgr.db.Create(&userinfo)
	return result.Error

}

func (mgr manager) IsExist(username string) (model.Userinfo, error) {
	var t_userinfo model.Userinfo
	result := mgr.db.Where("name=?", username).Find(&t_userinfo)
	return t_userinfo, result.Error
}
