package cache

import (
	"sync"
	"time"
)

type cacher struct {
	cache interface{}
	t     int //到期倒计时 秒
}

var cachemap sync.Map

func init() {
	go timeout()
}

func timeout() {
	t := time.NewTicker(time.Second)
	for range t.C {
		cachemap.Range(
			func(k, v interface{}) bool {
				vcacher := v.(cacher)
				if vcacher.t > 0 {
					vcacher.t = vcacher.t - 1
					cachemap.Store(k, vcacher)
				}
				if vcacher.t == 0 {
					cachemap.Delete(k)
				}
				return true
			},
		)
	}
}
func Set(k string, v interface{}, t int) {
	cachemap.Store(k, cacher{
		cache: v,
		t:     t,
	})
}
func Get(k string) interface{} {
	cache, ok := cachemap.Load(k)
	if ok {
		return cache.(cacher).cache
	}
	return nil
}

func Pull(k string) interface{} {
	cache, ok := cachemap.LoadAndDelete(k)
	if ok {
		return cache.(cacher).cache
	}
	return nil
}

func Del(k string) {
	cachemap.Delete(k)
}
