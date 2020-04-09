package src

import "github.com/gin-gonic/gin"

// 缓存装饰器
func CacheDecorator(h gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		// redis判断
	}
}
