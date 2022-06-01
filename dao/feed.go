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
	if count <= 30 {
		result := mgr.db.Model(&model.Video{}).Order("created_at DESC").Preload("Author").Find(&videos)
		return videos, result.Error
	}
	// "created_at <= ?" : 按道理是 <= , 但是查不全，不知道为啥那边请求的latestTime有时返回的是最早的时间
	// "created_at >= ?" : 当latestTime表示的是当前时间，“>=” 就会查询为空
	result := mgr.db.Where("created_at >= ?", time.Unix(latestTime, 0).Format(timeLayout)).Order("created_at DESC").Preload("Author").Limit(30).Find(&videos)
	return videos, result.Error
}

// UserToVideo 一个人访问另一个人的视频，查询是否点赞
func (mgr manager) UserToVideo(userinfo model.Userinfo, video model.Video) bool {
	var favorite model.Favorite
	if err := mgr.db.Model(model.Favorite{}).Where("user_id = ? AND video_id = ?", userinfo.Id, video.Id).Find(
		&favorite).Error; err == nil {
		return true
	} else {
		return false
	}

}
