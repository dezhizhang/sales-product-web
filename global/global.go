package global

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc"
	"sales-product-web/model"
)

var (
	Conn         *grpc.ClientConn
	Trans        ut.Translator
	NacosConfig  *model.NacosConfig  = &model.NacosConfig{}
	ServerConfig *model.ServerConfig = &model.ServerConfig{}
)
