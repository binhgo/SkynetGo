package cache

import (
	"time"
)

// MEM CACHE
type MemCache struct {
	Config     *Config
	cacheSpace string
}

func (mc *MemCache) Run() {
	go mc.loadDataToCache()
}

func (mc *MemCache) GetData() (string, error) {
	return mc.cacheSpace, nil
}

func (mc *MemCache) loadDataToCache() {
	for {
		// query data some where from database
		data := mc.Config.QueryFunction()

		// assign queried data to cache space
		mc.cacheSpace = data

		// sleep when doing nothing
		time.Sleep(time.Second * 60 * mc.Config.PeriodicInMinute)
	}
}
