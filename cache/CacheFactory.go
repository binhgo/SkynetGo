package cache

import (
	"errors"
	"log"
	"sync"
)

// Factory to create and manage caches
// ThresholdInKilobytes: size of string to decide save to memory or file
type CacheFactory struct {
	ThresholdInKilobytes int
	Caches               map[string]Cache
	Lock                 *sync.RWMutex
}

func (fac *CacheFactory) newCache(config *Config) (Cache, error) {

	// new Caches map
	if fac.Caches == nil {
		fac.Caches = make(map[string]Cache)
	}

	// check if cache exist, return
	// compare cache name, if exist, return error
	fac.Lock.RLock()
	_, ok := fac.Caches[config.CacheName]
	fac.Lock.RUnlock()
	if ok {
		return nil, errors.New("Cache with name: " + config.CacheName + " exist")
	}

	// end check

	var C Cache
	// check data size
	data := config.QueryFunction()
	sizeInByte := checkDataSize(data)
	sizeInKB := sizeInByte / 1024

	// depend on size, factory will create suitable cache type: memcache or filecache
	if sizeInKB > fac.ThresholdInKilobytes {
		// File cache
		C = &FileCache{config}
	} else {
		// Mem cache
		C = &MemCache{Config: config}
	}

	fac.Lock.Lock()
	fac.Caches[config.CacheName] = C
	fac.Lock.Unlock()

	return C, nil
}

func (fac *CacheFactory) NewCacheAndRun(config *Config) Cache {
	cache, err := fac.newCache(config)
	if err != nil {
		log.Fatalln(err)
	}

	cache.Run()
	return cache
}

func (fac *CacheFactory) GetData(cacheName string) (string, error) {

	fac.Lock.RLock()
	data, ok := fac.Caches[cacheName]
	fac.Lock.RUnlock()

	if !ok {
		return "", errors.New("cacheName does not exist")
	}

	result, err := data.GetData()
	if err != nil {
		return "", err
	}

	return result, nil
}
