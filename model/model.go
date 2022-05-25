package model

import "gorm.io/gorm"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64    `json:"id,omitempty"`
	Author        Userinfo `json:"author"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
}

type Comment struct {
	gorm.Model
	Id         int64    `gorm:"column:id;" json:"id,omitempty"`
	Commenter  Userinfo `gorm:"embedded;embeddedPrefix:user_" json:"user,omitempty"`
	Video      Video    `gorm:"embedded;embeddedPrefix:video_"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}

type Userinfo struct {
	gorm.Model
	Id            int64  `gorm:"primary_key;column:id;type:int;not null" json:"id,omitempty"`
	Name          string `gorm:"column:name;type:varchar;not null" json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
