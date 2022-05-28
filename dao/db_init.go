package dao

import (
	"douyin/model"
	"douyin/pkg/constrant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type manager struct {
	db *gorm.DB
}

type Manager interface {
	// Register 注册
	Register(user model.Userinfo) error
	// IsExist 判断是否用户是否已存在
	IsExist(username string) (model.Userinfo, error)

	// CommentAction 添加评论
	CommentAction(comment model.Comment, actionType string) error
	// GetUserInfo 获取用户信息(只返回响应需要的信息)
	GetUserInfo(userId int64) (model.Userinfo, error)
	// GetCommentList 获取评论列表
	GetCommentList(videoId int64) ([]model.Comment, error)

	// FavoriteAction 点赞操作
	FavoriteAction(userId int64, videoId int64, actionType string) error
	// FavoriteList 获取用户点赞列表
	FavoriteList(userId int64) ([]model.Video, error)
}

var Mgr Manager

func init() {
	db, err := gorm.Open(mysql.Open(constrant.MySQLDefaultDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled

		},
	})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}

}
