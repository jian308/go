package conf

import (
	"github.com/pelletier/go-toml"
)

var cfgtoml = "conf.toml"
var Data *toml.Tree

func Auto() {
	Load(cfgtoml)
}
func Load(cfgfile string) {
	config, err := toml.LoadFile(cfgfile) //加载toml文件
	if err != nil {
		panic(err)
	}
	Data = config
}

func Get(k string) interface{} {
	if Data == nil {
		return nil
	}
	return Data.Get(k)
}
