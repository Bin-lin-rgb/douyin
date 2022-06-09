package controller

import (
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	model.Response
	UserList []model.Userinfo `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserID := c.Query("to_user_id")
	toUser, _ := strconv.ParseInt(toUserID, 10, 64)
	actionType := c.Query("action_type")
	if user, exist := TokenIsValid(token); exist {

		err := dao.Mgr.RelationAction(user.Id, toUser, actionType)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "关注操作失败，请重试"})
			return
		}

		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FollowList(c *gin.Context) {
	token := c.Query("token")

	if user, exist := TokenIsValid(token); exist {
		uid := user.Id
		followList, err := dao.Mgr.GetFollowList(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "获取关注列表失败，请重试"})
			return
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: followList,
		})

	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}

func FollowerList(c *gin.Context) {
	token := c.Query("token")

	if user, exist := TokenIsValid(token); exist {
		uid := user.Id
		followList, err := dao.Mgr.GetFollowerList(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "获取粉丝列表失败，请重试"})
			return
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: followList,
		})

	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
