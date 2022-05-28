package dao

import (
	"douyin/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (mgr manager) FavoriteAction(userId int64, videoId int64, actionType string) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	var result *gorm.DB

	switch actionType {
	case "1":
		result = mgr.db.Create(&favorite)
	case "2":
		result = mgr.db.Where("user_id = ? AND video_id= ?", userId, videoId).Delete(&favorite)
	default:
		return errors.New("未知操作类型")
	}

	return result.Error

}

func (mgr manager) FavoriteList(userId int64) ([]model.Video, error) {

	var favoritedVideo []model.Favorite
	var favoriteList []model.Video

	tx := mgr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if result := tx.Where("user_id = ?", userId).Find(&favoritedVideo); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var videoId []int64
	for _, v := range favoritedVideo {
		videoId = append(videoId, v.VideoId)
	}

	fmt.Printf("%v\n", videoId)

	if result := tx.Preload("Author").Find(&favoriteList, videoId); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	return favoriteList, tx.Commit().Error
}
