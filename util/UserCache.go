package util

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var userCache *cache.Cache

func UserCache() *cache.Cache {
	if userCache == nil {
		userCache = cache.New(5*time.Minute, 10*time.Minute)
	}

	return userCache
}
