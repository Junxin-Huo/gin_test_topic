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
	topicId := c.Param("topic_id")
	topics := Topic{}
	DBHelper.Find(&topics, topicId)
	c.Set("dbResult", topics)

	//conn := RedisDefaultPool.Get()
	//defer conn.Close()
	//redisKey := "topic_" + topicId
	//res, err := redis.Bytes(conn.Do("get", redisKey))
	//if err != nil {
	//	// 缓存里没有值
	//	DBHelper.Find(&topics, topicId)
	//	bytes, err2 := json.Marshal(topics)
	//
	//	var redisTime int
	//	if topics.TopicID == 0 {
	//		// 从数据库中未匹配到数据
	//		redisTime = 20
	//	} else {
	//		// 正常数据
	//		redisTime = 60
	//	}
	//
	//	if err2 != nil {
	//		c.JSON(http.StatusInternalServerError, "json序列化错误，未缓存到redis")
	//		return
	//	} else {
	//		_, err2 := conn.Do("setex", redisKey, redisTime, bytes)
	//		if err2 != nil {
	//			log.Println("缓存redis失败", err2.Error())
	//		} else {
	//			log.Println("缓存redis成功")
	//		}
	//	}
	//	c.JSON(http.StatusOK, topics)
	//	log.Println("从数据库中读取")
	//	return
	//} else {
	//	// 缓存里有值
	//	if err3 := json.Unmarshal(res, &topics); err3 != nil {
	//		c.JSON(http.StatusInternalServerError, "redis中json解析错误")
	//		return
	//	}
	//	c.JSON(http.StatusOK, topics)
	//	log.Println("从redis中读取")
	//}
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
