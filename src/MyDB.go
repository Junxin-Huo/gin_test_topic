package src

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DBHelper *gorm.DB
var err error

func InitDB() {
	DBHelper, err = gorm.Open("mysql", "root:123456@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		//log.Fatal("初始化DB错误:", err)
		//panic(err)
		ShutDownServer(err)
		return
	}
	//DBHelper.LogMode(true)
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)
}
