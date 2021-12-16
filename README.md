# go
自用go的简易库
```
go get github.com/jian308/go
```
## 介绍
目前包含cache/conf/log可以使用
使用前需要引入对应的
```
import(
	"github.com/jian308/go/cache"
	"github.com/jian308/go/conf"
	"github.com/jian308/go/log"
)
```
### cache使用
```
//设置key=value缓存10秒 时间填-1永久缓存
cache.Set("key","value",10) 
//获取key的值 获取不到返回nil
cache.Get("key")
//删除key
cache.Del("key")
//获取的同时删掉key
cache.Pull("key")
```
### conf使用
```
//先要初始化 可以选择自动初始化
conf.Auto() //自动加载根目录conf.toml配置文件
//也可以自定义初始化对应tmol配置文件
conf.Load("./config/conf.toml")
//获取配置值 获取不到返回nil
conf.Get(key)
//获取到值之后也可以直接解析出对应格式
conf.Get(key).(string) //获取到string格式 对应toml文件里的格式
//如只有一级的用 conf.Get("appname") 二级用 conf.Get("mysql.host")
```
### log使用
```
//直接使用
log.Debug("test")
log.Debugf("test %d",123)
log.Info("test")
log.Infof("test %d",123)
log.Warn("test")
log.Warnf("test %d",123)
log.Error("test")
log.Errorf("test %d",123)
log.Fatal("test")
log.Fatalf("test %d",123)
//设置级别 Debug=-1 Info=0 Warn=1 Error=2 ...
//需要先定义级别
log.SetLevel(0) //只显示Info+Warn+...的信息
//存文件日志(会自动切割)
//需要先定义存日志的目录 默认保存到对应目录的log.log里
log.SetPath("./logs/")
//保存到./logs/log.log里
```