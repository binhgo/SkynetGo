package cache

import (
	"log"
	"sync"
	"testing"
	"time"

	"SkynetGo/util"
	"github.com/stretchr/testify/assert"
)

func TestCacheFactory_NewCache(t *testing.T) {
	ass := assert.New(t)

	lock := &sync.RWMutex{}

	cacheFactory := &CacheFactory{ThresholdInKilobytes: 1, Lock: lock}

	// test memcache
	config1 := &Config{"persons", 10, queryPersonFromMySQLTestMemCache}
	cache := cacheFactory.NewCacheAndRun(config1)
	// test
	ass.NotNil(cache)
	ass.Implements((*Cache)(nil), cache)
	_, ok := cache.(*MemCache)
	ass.True(ok)

	result, err := cacheFactory.GetData("person")
	ass.NotNil(err)
	log.Println(result)

	result, err = cacheFactory.GetData("persons")
	ass.Nil(err)
	log.Println(result)

	// test filecache
	config2 := &Config{"cities", 20, queryCityFromMySQLTestFileCache}
	fileCache := cacheFactory.NewCacheAndRun(config2)
	// test
	ass.NotNil(fileCache)
	ass.Implements((*Cache)(nil), fileCache)
	_, ok2 := fileCache.(*FileCache)
	ass.True(ok2)

	time.Sleep(time.Second * 1)
	result2, err := cacheFactory.GetData("cities")
	ass.Nil(err)
	log.Println(result2)
}

func queryCityFromMySQLTestFileCache() string {
	data, _ := util.ReadFile("./test.txt")
	return data
}

func queryPersonFromMySQLTestMemCache() string {
	return "This is test data from mysql"
}
