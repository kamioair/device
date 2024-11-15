package config

import (
	"github.com/kamioair/qf/utils/qconfig"
)

// Config 自定义配置
var Config = struct {
	IdLength int
	IdPrefix string
	StartId  uint64
}{
	IdLength: 6,
	IdPrefix: "",
	StartId:  100000,
}

func Init(module string) {
	qconfig.Load(module, &Config)
}
