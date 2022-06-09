package dao

import (
	"douyin/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func (mgr manager) RelationAction(userID, toUserID int64, actionType string) error {
	follow := model.Follow{
		UserId:   userID,
		FollowId: toUserID,
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
		{
			//更新follow表
			if err := tx.Create(&follow).Error; err != nil {
				log.Println(err)
				tx.Rollback()
				return err
			}

			//更新关注者的关注人数
			if err := tx.Model(&model.Userinfo{}).Where("id = ?", follow.UserId).
				UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
				log.Println(err)
				tx.Rollback()
				return err
			}

			//更新被关注者的粉丝人数
			if err := tx.Model(&model.Userinfo{}).Where("id = ?", follow.FollowId).
				UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
				log.Println(err)
				tx.Rollback()
				return err
			}

			return tx.Commit().Error
		}

	case "2":
		//删除follow表的字段
		if err := tx.Where("user_id = ? AND follow_id= ?", follow.UserId, follow.FollowId).Delete(&follow).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}

		//更新关注者的关注人数
		if err := tx.Model(&model.Userinfo{}).Where("id = ?", follow.UserId).
			UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}

		//更新被关注者的粉丝人数
		if err := tx.Model(&model.Userinfo{}).Where("id = ?", follow.FollowId).
			UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	default:
		return errors.New("未知操作类型")
	}
}

func (mgr manager) GetFollowList(userID int64) ([]model.Userinfo, error) {

	var followUserList []model.Follow
	var followList []model.Userinfo

	//找到该用户关注的用户列表
	if err := mgr.db.Where("user_id = ?", userID).Preload("User").Preload("Follow").Find(&followUserList).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("%v", followUserList)

	//关注用户的具体用户信息
	for _, v := range followUserList {
		//fmt.Println(v)
		v.Follow.IsFollow, _ = mgr.IsFollow(userID, v.Follow.Id)
		followList = append(followList, v.Follow)
	}

	return followList, nil
}

func (mgr manager) GetFollowerList(userID int64) ([]model.Userinfo, error) {
	var followUserList []model.Follow
	var followList []model.Userinfo

	//关注该用户的用户
	if err := mgr.db.Where("follow_id = ?", userID).Preload("User").Preload("Follow").Find(&followUserList).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("%v", followUserList)

	//粉丝列表的具体信息
	for _, v := range followUserList {
		//fmt.Println(v)
		v.User.IsFollow, _ = mgr.IsFollow(userID, v.User.Id)
		followList = append(followList, v.User)
	}

	return followList, nil
}

func (mgr manager) IsFollow(userId, toUserId int64) (bool, error) {
	var count int64
	var follow model.Follow
	if err := mgr.db.Where("user_id = ? AND follow_id = ?", userId, toUserId).
		Find(&follow).Count(&count).Error; err != nil {
		log.Println(err)
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
