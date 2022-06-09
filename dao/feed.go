package dao

import (
	"douyin/model"
	"fmt"
	"time"
)

func (mgr manager) GetAllVideo(latestTime int64) ([]model.Video, error) {
	var videos []model.Video
	timeLayout := "2006-01-02 15:04:05"
	var count int64
	mgr.db.Model(&model.Video{}).Count(&count)
	fmt.Println("---Count---", count)
	// "created_at <= ?" : 按道理是 <= , 但是查不全，不知道为啥那边请求的latestTime有时返回的是最早的时间
	// "created_at >= ?" : 当latestTime表示的是当前时间，“>=” 就会查询为空
	result := mgr.db.Model(&model.Video{}).Where("created_at <= ?", time.Unix(latestTime, 0).Format(timeLayout)).
		Order("created_at DESC").Preload("Author").Limit(30).Count(&count).Find(&videos)

	if count == 1 {
		//	证明已经刷完了，是小于等于最早时间只能查到1个视频记录
		result := mgr.db.Model(&model.Video{}).Where("created_at >= ?", time.Unix(latestTime, 0).Format(timeLayout)).
			Order("created_at DESC").Preload("Author").Limit(30).Count(&count).Find(&videos)
		return videos, result.Error
	}
	return videos, result.Error
}

// IsFavorite 一个人访问另一个人的视频，查询是否点赞
func (mgr manager) IsFavorite(userID int64, videoId int64) (bool, error) {
	var favorite model.Favorite
	var count int64

	if err := mgr.db.Where("user_id = ? AND video_id = ?", userID, videoId).
		Find(&favorite).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
