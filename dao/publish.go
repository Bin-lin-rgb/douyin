package dao

import "douyin/model"

func (mgr manager) InsertToVideo(video model.Video) error {
	result := mgr.db.Create(&video)
	return result.Error

}
