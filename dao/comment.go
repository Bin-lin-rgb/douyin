package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
)

func (mgr manager) CommentAction() error {
	return nil
}

func (mgr manager) AddComment(comment model.Comment) error {

	result := mgr.db.Create(&comment)
	return result.Error
}
