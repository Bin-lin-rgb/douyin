package model

import "gorm.io/gorm"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64    `json:"id,omitempty"`
	AuthorId      int64    `json:",omitempty" gorm:"column:author_id;"`
	Author        Userinfo `json:"author" gorm:"foreignKey:AuthorId;"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
}

type Comment struct {
	gorm.Model
	Id          int64    `json:"id,omitempty" gorm:"column:id;" `
	CommenterId int64    `json:",omitempty" gorm:"column:user_id;"`
	VideoId     int64    `json:",omitempty" gorm:"column:video_id;"`
	Content     string   `json:"content,omitempty" gorm:"column:content;"`
	CreateDate  string   `json:"create_date,omitempty" gorm:"-" `
	Commenter   Userinfo `json:"user,omitempty" gorm:"foreignKey:CommenterId;-;"`
	Video       Video    `json:"video,omitempty" gorm:"foreignKey:VideoId;-;"`
}

type Userinfo struct {
	gorm.Model
	Id            int64  `json:"id,omitempty" gorm:"primary_key;column:id;type:int;not null" `
	Name          string `json:"name,omitempty" gorm:"column:name;type:varchar;not null" `
	Password      string `json:"password,omitempty" gorm:"omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"omitempty"`
}

type Favorite struct {
	UserId  int64    `json:"user_id,omitempty" gorm:"column:user_id"`
	VideoId int64    `json:"video_id,omitempty" gorm:"column:video_id"`
	User    Userinfo `json:"user,omitempty" gorm:"foreignKey:UserId;-;"`
	Video   Video    `json:"video,omitempty" gorm:"foreignKey:VideoId;-;"`
}
