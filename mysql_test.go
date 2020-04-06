package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	hjxSrc "hjx.test/topic/src"
	"testing"
	"time"
)

func TestMysqlRaw(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Raw("select topic_id, topic_title from topics").Rows()
	for rows.Next() {
		var t_id int
		var t_title string
		rows.Scan(&t_id, &t_title)
		fmt.Println(t_id, t_title)
	}
}

func TestMysqlFrame(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	db.LogMode(true)
	// 自动匹配表名，不额外添加复数
	//db.SingularTable(true)

	//tc := hjxSrc.TopicClass{}
	//db.Table("topic_class").First(&tc, 2)
	//fmt.Println(tc)

	//var tcs []hjxSrc.TopicClass
	//db.Table("topic_class").Find(&tcs)
	//fmt.Println(tcs)

	//var tcs []hjxSrc.TopicClass
	//db.Table("topic_class").Where("class_name=?", "技术类").Find(&tcs)
	//fmt.Println(tcs)

	var tcs []hjxSrc.TopicClass
	db.SingularTable(true)
	db.Table("topic_class").Where(&hjxSrc.TopicClass{ClassName: "技术类"}).Find(&tcs)
	fmt.Println(tcs)
}

func TestNewTable(t *testing.T)  {
	db, err := gorm.Open("mysql", "root:123456@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	//db.LogMode(true)

	fmt.Println(db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&hjxSrc.Topic{}).RowsAffected)
	fmt.Println(db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&hjxSrc.Topic{}).RowsAffected)
}

func TestAdd(t *testing.T)  {
	db, err := gorm.Open("mysql", "root:123456@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	db.LogMode(true)

	topic := hjxSrc.Topic{
		TopicTitle:      "TopicTitle",
		TopicShortTitle: "TopicShortTitle",
		UserIP:          "127.0.0.1",
		TopicUrl:        "testurl",
		TopicDate:       time.Now(),
	}
	fmt.Println(db.Create(&topic).RowsAffected)
	fmt.Println(topic.TopicID)
}
