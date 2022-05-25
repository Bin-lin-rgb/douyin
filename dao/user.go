package dao

//数据库相关，gorm相关接口
import (
	"douyin/model"
)

type Manager interface {
	Register(user model.Userinfo) error
	IsExist(username string) (model.Userinfo, error)
	//GetPwd(pwd string) ([]byte, error)
	//ComparePwd(pwd1 string, pwd2 string) bool
}

//type manager struct {
//	db *gorm.DB
//}
//
//var Mgr Manager

//func init() {
//	db, err := gorm.Open(mysql.Open(constrant.MySQLDefaultDSN), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
//			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
//
//		},
//	})
//	if err != nil {
//		log.Fatal("Failed to init db:", err)
//	}
//	Mgr = &manager{db: db}
//	//db.AutoMigrate(&model.User{})
//
//}

func (mgr manager) Register(userinfo model.Userinfo) error {
	result := mgr.db.Create(&userinfo)
	return result.Error

}

func (mgr manager) IsExist(username string) (model.Userinfo, error) {
	var t_userinfo model.Userinfo
	result := mgr.db.Where("name=?", username).Find(&t_userinfo)
	return t_userinfo, result.Error
}
