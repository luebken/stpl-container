package analysis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

var RedisOptions = &redis.Options{
	Addr:     "",
	Password: "", // no password set
	DB:       0,  // use default DB
}

var API_KEY = ""

type ComponentInfo struct {
	Raw     string
	UsedAPI bool
}

func GetComponentInfo(platform string, groupid string, artefact string) ComponentInfo {

	result := ""
	usedAPI := false
	cachekey := "component" + ":" + platform + ":" + groupid + ":" + artefact

	log.Info("Getting Component info for " + cachekey)

	redisClient := redis.NewClient(RedisOptions)
	result, err := redisClient.Get(cachekey).Result()
	if err == redis.Nil {
		log.Info("Cache miss for key: ", cachekey)
	} else if err != nil {
		log.Errorf("Err when getting key %v. Error:%v ", cachekey, err)
	}

	if result == "" { // not cached
		log.Info("Not cached. Quering libraries.io.")
		request := fmt.Sprintf("https://libraries.io/api/%v/%v%%3A%v?api_key=%v", platform, groupid, artefact, API_KEY)
		resp, err := http.Get(request)
		usedAPI = true
		if err != nil {
			log.Errorf("Error calling http.Get %v", err)
			return ComponentInfo{result, usedAPI}

		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
			return ComponentInfo{result, usedAPI}
		}
		result = string(body)

		//cache result
		pipe := redisClient.TxPipeline()
		pipe.Set(cachekey, result, 0)
		pipe.Set(cachekey+":modified", time.Now().Unix(), 0)
		_, err = pipe.Exec()
		if err != nil {
			log.Error(err)
			return ComponentInfo{result, usedAPI}
		}

	} else {
		log.Info("Using cached result.")
	}

	return ComponentInfo{result, usedAPI}
}

func TestRedis() (bool, error) {
	redisClient := redis.NewClient(RedisOptions)
	pong, err := redisClient.Ping().Result()
	return pong == "PONG", err
}

func GetNumberOfCachedComponents() (int, error) {
	redisClient := redis.NewClient(RedisOptions)
	keys, err := redisClient.Keys("component*").Result()
	return len(keys) / 2, err //ignoring :modified keys
}
