package dao

import (
	"douyin/model"
	"douyin/pkg/constrant"
)

func (mgr manager) InsertToVideo(video model.Video) error {
	result := mgr.db.Create(&video)
	return result.Error

}

func (mgr manager) GetVideoByUserId(userId string) ([]model.Video, error) {
	var videos []model.Video
	result := mgr.db.Where("author_id=?", userId).Order("created_at DESC").Find(&videos)
	videoLen := len(videos)
	for i := 0; i < videoLen; i++ {
		videos[i].PlayUrl = StrBuilder(constrant.Root, videos[i].PlayUrl)
		videos[i].CoverUrl = StrBuilder(constrant.Root, videos[i].CoverUrl)
	}
	return videos, result.Error
}
