package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sales-product-web/global"
	"sales-product-web/model"
	"sales-product-web/proto"
	"sales-product-web/utils"
)

// GetBannerList 获取轮播图列表
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
	fmt.Println("rsp", rsp)

	utils.ResponseSuccess(c, &model.ResponseData{
		Code:  http.StatusOK,
		Msg:   "获取轮播图列表成功",
		Data:  rsp.Data,
		Total: rsp.Total,
	})
}

// CreateBanner 添加轮播图
func CreateBanner(c *gin.Context) {
	var banner model.Banner
	err := c.ShouldBindJSON(&banner)
	if err != nil {
		zap.S().Errorf("序列化失败%s", err.Error())
		return
	}
	rsp, err := global.BannerSrvClient.CreateBanner(context.Background(), &proto.CreateBannerRequest{
		Name:        banner.Name,
		Link:        banner.Link,
		Url:         banner.Url,
		Status:      banner.Status,
		Position:    banner.Position,
		Description: banner.Description,
	})
	if err != nil {
		zap.S().Errorw("新增轮播图失败%s", err.Error())
		return
	}
	utils.ResponseSuccess(c, &model.ResponseData{
		Code: http.StatusOK,
		Msg:  "新增轮播图成功",
		Data: rsp,
	})
}

// UpdateBanner 更新轮播图
func UpdateBanner(c *gin.Context) {
	var banner model.Banner
	err := c.ShouldBindJSON(&banner)
	if err != nil {
		zap.S().Errorf("序列化失败%s", err.Error())
		return
	}
	_, err = global.BannerSrvClient.UpdateBanner(context.Background(), &proto.UpdateBannerRequest{
		Id:          banner.Id,
		Name:        banner.Name,
		Link:        banner.Link,
		Url:         banner.Url,
		Status:      banner.Status,
		Position:    banner.Position,
		Description: banner.Description,
	})
	if err != nil {
		zap.S().Errorw("更新轮播图失败%s", err.Error())
		return
	}
	utils.ResponseSuccess(c, &model.ResponseData{
		Code: http.StatusOK,
		Msg:  "更新轮播图成功",
		Data: nil,
	})
}
