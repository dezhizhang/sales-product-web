package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sales-product-web/global"
	"sales-product-web/proto"
	"sales-product-web/utils"
)

func GetBannerList(c *gin.Context) {
	pageIndex := utils.StringToNumber(c.DefaultQuery("pageIndex", "1"))
	pageSize := utils.StringToNumber(c.DefaultQuery("pageSize", "10"))
	rsp, err := global.BannerSrvClient.GetBannerList(context.Background(), &proto.BannerRequest{
		PageIndex: int32(pageIndex),
		PageSize:  int32(pageSize),
	})
	if err != nil {
		zap.S().Errorw("获取轮播图列表失败%s", err.Error())
	}

	utils.ResponseSuccess(c, http.StatusOK, "获取轮播图成功", rsp.Data, rsp.Total)
}
