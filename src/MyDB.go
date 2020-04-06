package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DBHelper *gorm.DB
var err error

func init() {
	DBHelper, err = gorm.Open("mysql", "root:123456@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	DBHelper.LogMode(true)
}
