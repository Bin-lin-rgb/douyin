package controller

import (
	"douyin/dao"
	"douyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")
	user, _ := TokenIsValid(token)

	// 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	latestTimeString := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(latestTimeString, 10, 64)

	videos, err := dao.Mgr.GetAllVideo(latestTime)
	if err != nil {
		log.Println("GetAllVideo:", err)
	}

	videoLen := len(videos)
	fmt.Println("--------Number of video records returned-------", len(videos))
	if videoLen == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  model.Response{StatusCode: 0},
			VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
		return
	}

	// 如果未登录
	if user != (model.Userinfo{}) {
		for i := 0; i < videoLen-1; i++ {
			b := dao.Mgr.UserToVideo(user, videos[i])
			if !b {
				videos[i].IsFavorite = false
				continue
			}
			videos[i].IsFavorite = true
		}
	}

	// 本次结构体数组最后一个的下标，
	earliestTime := videoLen - 1
	// NextTime 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  videos[earliestTime].Model.CreatedAt.Unix(),
	})

}
