package global

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc"
	"sales-product-web/model"
	"sales-product-web/proto"
)

var (
	Conn             *grpc.ClientConn
	Trans            ut.Translator
	ProductSrvClient proto.ProductClient
	NacosConfig      *model.NacosConfig  = &model.NacosConfig{}
	ServerConfig     *model.ServerConfig = &model.ServerConfig{}
)
