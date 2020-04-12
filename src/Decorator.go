package src

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

// 缓存装饰器
func CacheDecorator(h gin.HandlerFunc, param string,
	redisKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// redis判断
		getID := context.Param(param)
		redisKey := fmt.Sprintf(redisKeyPattern, getID)

		conn := RedisDefaultPool.Get()
		defer conn.Close()
		res, err := redis.Bytes(conn.Do("get", redisKey))
		if err != nil {
			// 缓存里没有值
			h(context) // 执行目标方法
			dbResult, exists := context.Get("dbResult")
			if !exists {
				dbResult = empty
			}
			bytes, err2 := json.Marshal(dbResult)

			var redisTime int
			topic := dbResult.(Topic)
			if topic.TopicID == 0 {
				// 从数据库中未匹配到数据
				redisTime = 20
			} else {
				// 正常数据
				redisTime = 60
			}

			if err2 != nil {
				context.JSON(http.StatusInternalServerError, "json序列化错误，未缓存到redis")
				return
			} else {
				_, err2 := conn.Do("setex", redisKey, redisTime, bytes)
				if err2 != nil {
					log.Println("缓存redis失败", err2.Error())
				} else {
					log.Println("缓存redis成功")
				}
			}
			context.JSON(http.StatusOK, dbResult)
			log.Println("从数据库中读取")
		} else {
			// 缓存里有值
			if err3 := json.Unmarshal(res, &empty); err3 != nil {
				context.JSON(http.StatusInternalServerError, "redis中json解析错误")
				return
			}
			context.JSON(http.StatusOK, empty)
			log.Println("从redis中读取")
		}
	}
}
