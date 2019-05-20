package cache

import (
	"log"
	"time"

	"SkynetGo/util"
)

// FILE CACHE
type FileCache struct {
	Config *Config
}

func (fc *FileCache) Run() {
	go fc.loadDataToCache()
}

func (fc *FileCache) GetData() (string, error) {
	// read file
	filePath := "./" + fc.Config.CacheName
	data, err := util.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// return data
	return data, nil
}

func (fc *FileCache) loadDataToCache() {
	for {
		// query data some where from database
		data := fc.Config.QueryFunction()

		// store data into file
		filePath := "./" + fc.Config.CacheName
		err := util.WriteFile(filePath, data)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * 60 * fc.Config.PeriodicInMinute)
	}
}
