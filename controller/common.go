package controller

import (
	"douyin/dao"
	"douyin/model"
	"log"
)

func TokenIsValid(token string) (model.Userinfo, bool) {
	var userinfo model.Userinfo

	// 是否过期
	claims, err := ParseToken(token)
	if err != nil {
		return userinfo, false
	}

	user, err := dao.Mgr.IsExist(claims.Username)
	if err != nil {
		log.Println(err)
		return userinfo, false
	}

	if user.Name != claims.Username {
		return userinfo, false
	}

	return user, true
}
