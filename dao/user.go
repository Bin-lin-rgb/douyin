package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
)

func (mgr manager) Register(userinfo model.Userinfo) error {
	result := mgr.db.Create(&userinfo)
	return result.Error

}
