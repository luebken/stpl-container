package analysis

import (
	"errors"
	"net/url"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

var redisOptions redis.Options
var librariesIoAPIKey string

// GetEnvRedisURL gets REDIS_URL defaults to redis://localhost:6379
func GetEnvRedisURL() error {
	redisURLString := os.Getenv("REDIS_URL")
	if redisURLString == "" {
		log.Info("Didn't find env REDIS_URL. Using default redis://localhost:6379")
		redisURLString = "redis://localhost:6379"
	}
	log.Info("using redis url: " + redisURLString)
	redisURL, err := url.Parse(redisURLString)
	if err != nil {
		log.Error("couldn't parse url: " + redisURLString)
		return err
	}
	if redisURL.Host == "" {
		log.Error("couldn't parse: " + redisURLString)
		return errors.New("couldn't parse " + redisURLString)
	}

	redisOptions.Addr = redisURL.Host
	if redisURL.User != nil {
		redisOptions.Password, _ = redisURL.User.Password()
	}
	redisOptions.DB = 0

	ping, err := testRedisConnection()
	if err != nil || !ping {
		log.Error("Couldn't ping redis. Err: ", err)
		return err
	}
	return nil
}
func testRedisConnection() (bool, error) {
	redisClient := redis.NewClient(&redisOptions)
	pong, err := redisClient.Ping().Result()
	return pong == "PONG", err
}

// GetLibrariesIoAPIKey gets env LIBRARIES_IO_API_KEY
func GetLibrariesIoAPIKey() error {
	librariesIoAPIKey = os.Getenv("LIBRARIES_IO_API_KEY")
	if librariesIoAPIKey == "" {
		return errors.New("didn't find env LIBRARIES_IO_API_KEY")
	}
	return nil
}
