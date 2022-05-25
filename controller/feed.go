package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/model"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
