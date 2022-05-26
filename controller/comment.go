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

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := TokenIsValid(token); exist {
		fmt.Println(user)
		//发布评论
		if actionType == "1" {
			text := c.Query("comment_text")
			videoId := c.Query("video_id")

			uid := user.Id

			fmt.Println("userId", uid)
			fmt.Println("videoId", videoId)

			vid, err := strconv.ParseInt(videoId, 10, 64)
			if err != nil {
				log.Println(err)
				return
			}

			//插入视频实现后删除，仅测试用
			vid = 2

			comment := model.Comment{
				Content:     text,
				CommenterId: uid,
				VideoId:     vid,
			}
			err = dao.Mgr.AddComment(comment)
			if err != nil {
				log.Println(err)
				return
			}

			userinfo, err := dao.Mgr.GetUserInfo(uid)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%v\n", userinfo)

			c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0},
				Comment: model.Comment{
					Commenter:  userinfo,
					Content:    text,
					CreateDate: time.Now().Format("01-02"),
				}})
		}
		//c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
