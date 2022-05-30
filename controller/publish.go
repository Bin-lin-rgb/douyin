package controller

import (
	"douyin/dao"
	"douyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	user, v := TokenIsValid(token)
	if !v {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "token is not valid."})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

	// ====
	root := "http://10.0.2.2:8080/static/"
	playurl := StrBulider(root, finalName)
	fmt.Println("-------------------", playurl)
	video := model.Video{
		AuthorId:      user.Id,
		Title:         title,
		PlayUrl:       playurl,
		CoverUrl:      "https://img0.baidu.com/it/u=3346653715,2652099287&fm=253&fmt=auto&app=138&f=JPEG?w=775&h=500",
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err = dao.Mgr.InsertToVideo(video)
	if err != nil {
		fmt.Println("InsertToVideo error:", err)
	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")

	_, b := TokenIsValid(token)
	if !b {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Token is not valid."},
		})
		return
	}

	videos, err := dao.Mgr.GetVideoByUserId(userId)
	if err != nil {
		log.Println("ReturnVideoByUserId:", err)
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}

func StrBulider(first string, finalName string) string {
	var builder strings.Builder
	builder.WriteString(first)
	builder.WriteString(finalName)
	return builder.String()
}
