package startup

import (
	"time"

	"github.com/FarmerChillax/AirRocket/vars"
	"github.com/patrickmn/go-cache"
)

// SetupVars 加载变量
func SetupVars() error {
	// 初始化缓存
	cache := cache.New(time.Duration(vars.AppSettings.DurationTime), 10*time.Minute)
	arc := vars.AirRocketCache{Cache: cache}
	vars.Cache = &arc

	return nil
}
