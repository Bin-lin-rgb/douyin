package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
)

type commentManager interface {
	CommentAction(comment model.Comment) error
}

func (mgr manager) CommentAction(userinfo model.Userinfo) error {
	result := mgr.db.Create(&userinfo)
	return result.Error

}

