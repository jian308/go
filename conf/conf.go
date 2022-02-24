package conf

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

var cfgtoml = "conf.toml"
var Data *toml.Tree

func Fdir(f string) string {
	fdir := f
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err := os.Stat(fdir) //os.Stat获取文件信息
	if os.IsNotExist(err) {
		fdir = dir + "/" + f
	}
	return fdir
}

func Auto() {
	Load(cfgtoml)
}
func Load(cfgfile string) {
	cfgtoml = cfgfile
	config, err := toml.LoadFile(cfgfile) //加载toml文件
	if err != nil {
		panic(err)
	}
	Data = config
}

func Get(k string) interface{} {
	if Data == nil {
		Auto()
	}
	return Data.Get(k)
}
