package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var db *gorm.DB

func init(){
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	myDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db = myDb
}

func GetDb() *gorm.DB{
	return db
}
