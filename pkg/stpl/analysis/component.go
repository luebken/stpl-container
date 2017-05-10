package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

type ComponentInfo struct {
	LibrariesIoComponent LibrariesIoComponent
	UsedAPI              bool
}

func GetComponentInfo(platform string, dependency string) ComponentInfo {

	var libioc LibrariesIoComponent
	usedAPI := false
	cachekey := "component" + ":" + platform + ":" + dependency

	log.Info("Getting Component info for " + cachekey)

	redisClient := redis.NewClient(&redisOptions)
	jsonRaw, err := redisClient.Get(cachekey).Result()
	if err == redis.Nil {
		log.Info("Cache miss for key: ", cachekey)
	} else if err != nil {
		log.Errorf("Err when getting key %v. Error:%v ", cachekey, err)
	}
	request := fmt.Sprintf("https://libraries.io/api/%v/%v?api_key=%v", platform, dependency, librariesIoAPIKey)
	log.Println(request)

	if jsonRaw == "" { // not cached
		log.Info("Not cached. Quering libraries.io.")
		resp, err := http.Get(request)
		usedAPI = true
		if err != nil {
			log.Errorf("Error calling http.Get %v", err)
			return ComponentInfo{libioc, usedAPI}

		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		log.Infof("Response. Status: %v. Length: %v", resp.Status, len(body))

		if err != nil {
			log.Error(err)
			return ComponentInfo{libioc, usedAPI}
		}

		err = json.Unmarshal(body, &libioc)
		if err != nil {
			log.Errorf("Error unmarshaling Error: %v\nJson:\n %v", err, string(body))
			return ComponentInfo{libioc, usedAPI}
		}

		//cache result
		pipe := redisClient.TxPipeline()
		pipe.Set(cachekey, string(body), 0)
		pipe.Set(cachekey+":modified", time.Now().Unix(), 0)
		_, err = pipe.Exec()
		if err != nil {
			log.Error(err)
			return ComponentInfo{libioc, usedAPI}
		}
	} else {
		log.Info("Using cached result.")
		err = json.Unmarshal([]byte(jsonRaw), &libioc)
		if err != nil {
			log.Errorf("Error unmarshaling Error: %v\nJson:\n %v", err, jsonRaw)
			return ComponentInfo{libioc, usedAPI}
		}

	}

	return ComponentInfo{libioc, usedAPI}
}

func GetNumberOfCachedComponents() (int, error) {
	redisClient := redis.NewClient(&redisOptions)
	keys, err := redisClient.Keys("component*").Result()
	return len(keys) / 2, err //ignoring :modified keys
}
