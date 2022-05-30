package dao

import (
	"douyin/model"
	"errors"
	"fmt"
	"log"
)

func (mgr manager) FavoriteAction(userId int64, videoId int64, actionType string) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}

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
		if err := tx.Create(&favorite).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	case "2":
		if err := tx.Where("user_id = ? AND video_id= ?", userId, videoId).Delete(&favorite).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	default:
		return errors.New("未知操作类型")
	}

	var favoriteCount int64

	if err := tx.Model(&model.Favorite{}).Where("video_id = ?", videoId).Count(
		&favoriteCount).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.Video{}).Select("favorite_count").Where("id = ?", videoId).Updates(
		model.Video{FavoriteCount: favoriteCount}).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

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

	if err := tx.Where("user_id = ?", userId).Find(&favoritedVideo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var videoId []int64
	for _, v := range favoritedVideo {
		videoId = append(videoId, v.VideoId)
	}

	fmt.Printf("%v\n", videoId)

	if err := tx.Preload("Author").Find(&favoriteList, videoId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return favoriteList, tx.Commit().Error
}
