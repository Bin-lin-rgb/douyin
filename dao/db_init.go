package dao

import (
	"douyin/model"
	"douyin/pkg/constrant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type manager struct {
	db *gorm.DB
}

type Manager interface {
	Register(user model.Userinfo) error
	IsExist(username string) (model.Userinfo, error)
	//GetPwd(pwd string) ([]byte, error)
	//ComparePwd(pwd1 string, pwd2 string) bool
}

var Mgr Manager

func init() {
	db, err := gorm.Open(mysql.Open(constrant.MySQLDefaultDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled

		},
	})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	//db.AutoMigrate(&model.User{})

}
