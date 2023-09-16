package utils

import "github.com/gin-gonic/gin"

// ResponseSuccess 成功时返回的参数
func ResponseSuccess(c *gin.Context, code int, msg string, data interface{}, total int32) {
	c.JSON(code, gin.H{"code": code, "msg": msg, "data": data, "total": total})
}
