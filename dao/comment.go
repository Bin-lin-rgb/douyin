package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
	"errors"
	"log"
)

func (mgr manager) CommentAction(comment model.Comment, actionType string) error {

	tx := mgr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	switch actionType {
	case "1":
		if err := tx.Create(&comment).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	case "2":
		//var commentId string

		if err := tx.Model(model.Comment{}).Where("user_id=? AND video_id =?", comment.CommenterId, comment.VideoId).Select("id").Find(&comment).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}

		if err := tx.Delete(&comment, comment.Id).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	default:
		return errors.New("未知操作类型")
	}

	var commentCount int64

	if err := tx.Model(&model.Comment{}).Where("video_id = ?", comment.VideoId).Count(
		&commentCount).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.Video{}).Select("comment_count").Where("id = ?", comment.VideoId).Updates(
		model.Video{CommentCount: commentCount}).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (mgr manager) GetCommentList(videoId int64) ([]model.Comment, error) {

	var commentList []model.Comment

	result := mgr.db.Where("video_id = ?", videoId).Preload("Commenter").Find(&commentList)

	return commentList, result.Error
}
