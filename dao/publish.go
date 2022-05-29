package dao

import (
	"douyin/model"
)

func (mgr manager) InsertToVideo(video model.Video) error {
	result := mgr.db.Create(&video)
	return result.Error

}

func (mgr manager) GetVideoByUserId(userId string) ([]model.Video, error) {
	var video []model.Video
	result := mgr.db.Where("author_id=?", userId).Find(&video)
	return video, result.Error
}
