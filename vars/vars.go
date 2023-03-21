package vars

import "github.com/patrickmn/go-cache"

type AirRocketCache struct {
	*cache.Cache
}

var Cache *AirRocketCache
