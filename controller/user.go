package controller

import (
	"douyin/dao"
	"douyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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

	// 查找用户是否存在
	user, err := dao.Mgr.IsExist(username)
	if err != nil {
		log.Println(err)
	}

	if user.Name != "" {
		fmt.Println("已存在！")
		return
	}

	encrypted, _ := GetPwd(password)

	userinfo := model.Userinfo{
		Name:     username,
		Password: string(encrypted),
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

	// 查找用户是否存在
	user, err := dao.Mgr.IsExist(username)
	if err != nil {
		log.Println(err)
	}

	if user.Name == "" {
		fmt.Println("不存在", user)

	}

	if ComparePwd(user.Password, password) {
		fmt.Println("登陆成功！")
	} else {
		fmt.Println("密码错误！")
	}

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

func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
