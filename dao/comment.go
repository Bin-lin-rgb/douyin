package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
	"errors"
	"gorm.io/gorm"
)

func (mgr manager) CommentAction(comment model.Comment, actionType string) error {

	var result *gorm.DB

	switch actionType {
	case "1":
		result = mgr.db.Create(&comment)
	case "2":
		//没有这个功能
	default:
		return errors.New("未知操作类型")
	}

	return result.Error
}

func (mgr manager) GetCommentList(videoId int64) ([]model.Comment, error) {

	var commentList []model.Comment

	result := mgr.db.Where("video_id = ?", videoId).Preload("Commenter").Find(&commentList)

	return commentList, result.Error
}
