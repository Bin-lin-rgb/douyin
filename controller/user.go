package controller

import (
	"douyin/dao"
	"douyin/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//token := username + password
	token, _ := GenToken(username, password)

	// 查找用户是否存在
	user, err := dao.Mgr.IsExist(username)
	if err != nil {
		log.Println(err)
	}

	if user.Name != "" {
		fmt.Println("已存在！")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}

	// encrypted : 已加密的密码
	encrypted, _ := GetPwd(password)
	userinfo := model.Userinfo{
		Name:     username,
		Password: string(encrypted),
	}
	// 将加密的密码写入数据库
	err = dao.Mgr.Register(userinfo)
	if err != nil {
		log.Println(err)
	}

	atomic.AddInt64(&userIdSequence, 1)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   userinfo.Id,
		Token:    token,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, _ := GenToken(username, password)

	// 查找用户是否存在
	user, err := dao.Mgr.IsExist(username)
	if err != nil {
		log.Println(err)
	}

	if user.Name == "" {
		fmt.Println("用户不存在！")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}

	if ComparePwd(user.Password, password) {
		fmt.Println("登陆成功！")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Password is not correct"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//user_id := c.Query("user_id")

	user, isvalid := TokenIsValid(token)

	if !isvalid {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Token is not valid."},
		})
		return
	}

	//user.IsFollow = true
	user.Password = ""

	c.JSON(http.StatusOK, UserResponse{
		Response: model.Response{StatusCode: 0},
		User:     user,
	})

}

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
