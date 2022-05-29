package controller

import "douyin/model"

var DemoVideos = []model.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://10.0.2.2:8080/static/bear.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []model.Comment{
	{
		Id:         1,
		Commenter:  DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = model.Userinfo{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
