package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 必须登录
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "缺少token参数")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetTopicDetail(c *gin.Context) {
	if err != nil {
		c.String(http.StatusBadRequest, "topic_id字段错误:%s", c.Param("topic_id"))
	} else {
		topicId := c.Param("topic_id")
		topics := Topic{}
		DBHelper.Find(&topics, topicId)
		c.JSON(http.StatusOK, topics)
	}
}

// 单帖子新增
func NewTopic(c *gin.Context) {
	//判断登录
	topic := Topic{}
	if err := c.BindJSON(&topic); err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, topic)
	}
}

// 多帖子新增
func NewTopics(c *gin.Context) {
	//判断登录
	Topics := Topics{}
	if err := c.BindJSON(&Topics); err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, Topics)
	}
}

func DelTopic(c *gin.Context) {
	//判断登录
	c.String(http.StatusOK, "删除帖子")
}

func GetTopicList(c *gin.Context) {
	//if c.Query("username") == "" {
	//	c.String(http.StatusOK, "获取帖子列表")
	//} else {
	//	c.String(http.StatusOK, "获取topicid=%s的帖子", c.Query("username"))
	//}

	query := TopicQuery{}
	if err := c.BindQuery(&query); err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}
