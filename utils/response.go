package utils

import (
	"github.com/gin-gonic/gin"
	"sales-product-web/model"
)

// ResponseSuccess 成功时返回的参数
func ResponseSuccess(c *gin.Context, opt *model.ResponseData) {
	if opt.Total == 0 {
		c.JSON(opt.Code, gin.H{
			"code": opt.Code,
			"msg":  opt.Msg,
			"data": opt.Data,
		})
		return
	}
	c.JSON(opt.Code, gin.H{
		"code":  opt.Code,
		"msg":   opt.Msg,
		"data":  opt.Data,
		"total": opt.Total,
	})
}
