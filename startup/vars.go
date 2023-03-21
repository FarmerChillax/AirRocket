package startup

import (
	"time"

	"github.com/FarmerChillax/AirRocket/vars"
	"github.com/patrickmn/go-cache"
)

// SetupVars 加载变量
func SetupVars() error {
	// 初始化缓存
	cache := cache.New(time.Duration(vars.AppSettings.DurationTime)*time.Second, 10*time.Minute)
	arc := vars.AirRocketCache{Cache: cache}
	vars.Cache = &arc

	return nil
}

// 加载客户端 server 的变量
func SetupNodeServerVars() error {

	return nil
}
