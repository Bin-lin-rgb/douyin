package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
	"log"
	"time"
)

type CommentResponse struct {
	Id          int64  `json:"id,omitempty" gorm:"column:id;" `
	CommenterId int64  `json:",omitempty" gorm:"column:user_id;"`
	VideoId     int64  `json:",omitempty" gorm:"column:video_id;"`
	Content     string `json:"content,omitempty" gorm:"column:content;"`
	CreatedAt   time.Time
}

func (mgr manager) CommentAction() error {
	return nil
}

func (mgr manager) AddComment(comment model.Comment) error {

	result := mgr.db.Create(&comment)
	return result.Error
}

func (mgr manager) GetCommentList(videoId int64) ([]model.Comment, error) {

	var commentList []model.Comment

	result := mgr.db.Where("video_id = ?", videoId).Preload("Commenter").Find(&commentList)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return commentList, nil
}
