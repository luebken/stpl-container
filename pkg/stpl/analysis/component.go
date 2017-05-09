package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

var RedisOptions = &redis.Options{}

func SetRedisURL(rurl string) {
	log.Info("using redis url: " + rurl)
	redisURL, _ := url.Parse(rurl)
	RedisOptions.Addr = redisURL.Host
	if redisURL.User != nil {
		RedisOptions.Password, _ = redisURL.User.Password()
	}
	RedisOptions.DB = 0
}

var API_KEY = ""

type ComponentInfo struct {
	LibrariesIoComponent LibrariesIoComponent
	UsedAPI              bool
}

func GetComponentInfo(platform string, groupid string, artefact string) ComponentInfo {

	var libioc LibrariesIoComponent
	usedAPI := false
	cachekey := "component" + ":" + platform + ":" + groupid + ":" + artefact

	log.Info("Getting Component info for " + cachekey)

	redisClient := redis.NewClient(RedisOptions)
	jsonRaw, err := redisClient.Get(cachekey).Result()
	if err == redis.Nil {
		log.Info("Cache miss for key: ", cachekey)
	} else if err != nil {
		log.Errorf("Err when getting key %v. Error:%v ", cachekey, err)
	}
	request := fmt.Sprintf("https://libraries.io/api/%v/%v%%3A%v?api_key=%v", platform, groupid, artefact, API_KEY)
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
