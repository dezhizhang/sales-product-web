package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sales-product-web/global"
	"sales-product-web/proto"
)

// 初始化grpc链接

func InitSrvConn() {

	var err error
	productSrvAddress := ""

	//fmt.Println(global.ServerConfig)

	name := global.ServerConfig.ProductSrv.Name
	cfg := api.DefaultConfig()
	consulConfig := global.ServerConfig.ConsulConfig
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("fmt", fmt.Sprintf(`Service == "%s"`, name))
	data, err1 := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	fmt.Println("data", data)
	if err1 != nil {
		panic(err1)
	}

	for _, value := range data {
		fmt.Println("value", value)
		productSrvAddress = value.Address
		break
	}
	if productSrvAddress == "" {
		zap.S().Errorw("initSrvConn 连接用户服务失败", "msg", err.Error())
	}

	fmt.Println("userSrvAddress", productSrvAddress)

	productConn, err := grpc.Dial(productSrvAddress, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("initSrvConn 连接用户服务失败", "msg", err.Error())
	}

	// 商品服务连接
	productSrvClient := proto.NewProductClient(productConn)

	global.ProductSrvClient = productSrvClient

}
