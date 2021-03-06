package controller

import (
	"douyin/dao"
	"douyin/model"
	"douyin/pkg/constrant"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := TokenIsValid(token); exist {
		videoId := c.Query("video_id")
		vid, err := strconv.ParseInt(videoId, 10, 64)
		if err != nil {
			log.Println(err)
			return
		}

		uid := user.Id

		err = dao.Mgr.FavoriteAction(uid, vid, actionType)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "点赞操作失败，请重试"})
			return
		}

		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token")

	if user, exist := TokenIsValid(token); exist {
		uid := user.Id
		favoriteList, err := dao.Mgr.FavoriteList(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "获取点赞列表失败，请重试"})
			return
		}

		videoLen := len(favoriteList)
		for i := 0; i < videoLen; i++ {
			favoriteList[i].PlayUrl = StrBuilder(constrant.Root, favoriteList[i].PlayUrl)
			favoriteList[i].CoverUrl = StrBuilder(constrant.Root, favoriteList[i].CoverUrl)
		}

		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			VideoList: favoriteList,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}
