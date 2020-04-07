package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	hjxSrc "hjx.test/topic/src"
	"os"
	"os/signal"
	"time"
)

func main() {
	var count int
	go func() {
		count = 0
		for {
			fmt.Println("执行", count)
			count++
			time.Sleep(time.Second)
		}
	}()

	signals := make(chan os.Signal)

	go func() {
		fmt.Println("1")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		fmt.Println("2")
		select {
		case <-ctx.Done():
			fmt.Println(3)
			signals <- os.Interrupt
		}
	}()

	signal.Notify(signals, os.Interrupt)
	s := <-signals
	fmt.Println(s)
}

func main2() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", hjxSrc.TopicUrl)
		//验证长度
		v.RegisterValidation("topics", hjxSrc.TopicsValidate)
	}

	// 单条帖子
	v1 := router.Group("/v1/topics")
	{
		v1.GET("", hjxSrc.GetTopicList)

		//v1.GET("/:topic_id", func(c *gin.Context) {
		//	c.String(http.StatusOK, "topic=%s", c.Param("topic_id"))
		//})

		v1.GET("/:topic_id", hjxSrc.GetTopicDetail)

		v1.Use(hjxSrc.MustLogin())
		{
			v1.POST("", hjxSrc.NewTopic)
			v1.DELETE("/:topic_id", hjxSrc.DelTopic)
		}
	}

	// 多条帖子
	v2 := router.Group("/v1/mtopics")
	{
		v2.Use(hjxSrc.MustLogin())
		{
			v2.POST("", hjxSrc.NewTopics)
		}
	}

	router.Run()
}
