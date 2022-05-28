package controller

import (
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction 评论操作
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := TokenIsValid(token); exist {
		//发布评论

		text := c.Query("comment_text")
		videoId := c.Query("video_id")

		uid := user.Id

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
		err = dao.Mgr.CommentAction(comment, actionType)
		if err != nil {
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0},
			Comment: model.Comment{
				Commenter:  user,
				Content:    text,
				CreateDate: user.CreatedAt.Format("01-02"),
			}})

	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")

	vid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	if _, exist := TokenIsValid(token); exist {

		commentList, err := dao.Mgr.GetCommentList(vid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, CommentListResponse{
				Response: model.Response{StatusCode: 500, StatusMsg: "获取评论列表失败，请重试"},
			})
			return
		}

		c.JSON(http.StatusOK, CommentListResponse{
			Response:    model.Response{StatusCode: 0},
			CommentList: commentList,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}
