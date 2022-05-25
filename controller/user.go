package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo/dao"
	"simple-demo/model"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var usersLoginInfo = map[string]model.Userinfo{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}
var userIdSequence = int64(1)

// -------

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.Userinfo `json:"user"`
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	//usersLoginInfo := make(map[string]model.User)
	//usersLoginInfo[]=username

	//usersLoginInfo[username]=username

	user, err := dao.Mgr.IsExist(username)
	if err != nil {
		log.Println(err)
	}

	if user.Name != "" {
		fmt.Println("存在", user)
		return
	}

	userinfo := model.Userinfo{
		Name:     username,
		Password: password,
	}
	err = dao.Mgr.Register(userinfo)
	if err != nil {
		log.Println(err)
	}

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.Userinfo{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
