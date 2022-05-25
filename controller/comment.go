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

	if user, exist := usersLoginInfo[token]; exist {
		//发布评论
		if actionType == "1" {
			text := c.Query("comment_text")
			videoId := c.Query("video_id")
			userId := c.Query("user_id")

			fmt.Println("userId", userId)
			fmt.Println("videoId", videoId)

			uid, err := strconv.ParseInt(userId, 10, 64)
			if err != nil {
				log.Println(err)
				return
			}
			vid, err := strconv.ParseInt(videoId, 10, 64)
			if err != nil {
				log.Println(err)
				return
			}

			comment := model.Comment{
				Content:   text,
				Commenter: model.Userinfo{Id: uid},
				Video:     model.Video{Id: vid},
			}

			err = dao.Mgr.AddComment(comment)
			if err != nil {
				log.Println(err)
				return
			}

			userinfo := model.Userinfo{}

			err = dao.Mgr.GetUserInfo(uid, &userinfo)
			if err != nil {
				log.Println(err)
				return
			}

			c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0},
				Comment: model.Comment{
					Commenter:  user,
					Content:    text,
					CreateDate: time.Now().Format("01-02"),
				}})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
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
