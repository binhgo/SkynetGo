package cache

import (
	"time"
	"unsafe"
)

type Cache interface {
	// public
	GetData() (string, error)
	// public
	Run()
	// private
	loadDataToCache()
}

// config struct
// CacheName: name of the cache: can be "city", "person"...
// PeriodicInMinute: time in minute to query new data and put to cache
// QueryFunction: type FuncQueryData func() string: function to query data
type Config struct {
	CacheName        string
	PeriodicInMinute time.Duration
	QueryFunction    FuncQueryData
}

// func to query data
type FuncQueryData func() string

// check string size in byte
func checkDataSize(data string) int {
	stringSize := len(data) + int(unsafe.Sizeof(data))
	return stringSize
}
