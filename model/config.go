package model

type ProductSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type ServerConfig struct {
	Name         string           `mapstructure:"name" json:"name"`
	Host         string           `mapstructure:"host" json:"host"`
	Port         int              `mapstructure:"port" json:"port"`
	ConsulConfig ConsulConfig     `mapstructure:"consul" json:"consul"`
	ProductSrv   ProductSrvConfig `mapstructure:"sales_product_srv" json:"sales_product_srv"`
}
