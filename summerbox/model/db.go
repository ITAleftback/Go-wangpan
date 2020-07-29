package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//连接数据库============================================================================
var (
	DB *gorm.DB
)
func init(){
	mysql, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user?charset=utf8")

	if err != nil {
		panic(err)
	}

	DB=mysql
	if !DB.HasTable(&User{}) {
		DB.CreateTable(&User{})
	}
}