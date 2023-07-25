package session

import (
	"log"
	"time"

	"github.com/syamsv/apollo-server/config"
	"github.com/syamsv/apollo-server/pkg/cache"
)

var manager struct {
	Auth   *cache.Cache
	Api    *cache.Cache
	Verify *cache.Cache
}

func InitCacheManager() {
	authConfig := cache.RedisConfig{
		Host:         config.REDIS_HOST,
		Port:         config.REDIS_PORT,
		Password:     config.REDIS_PASSWORD,
		DB:           0,
		MaxRetries:   3,
		RetryBackoff: 2 * time.Second,
	}
	apiConfig := cache.RedisConfig{
		Host:         config.REDIS_HOST,
		Port:         config.REDIS_PORT,
		Password:     config.REDIS_PASSWORD,
		DB:           0,
		MaxRetries:   3,
		RetryBackoff: 2 * time.Second,
	}
	verifyConfig := cache.RedisConfig{
		Host:         config.REDIS_HOST,
		Port:         config.REDIS_PORT,
		Password:     config.REDIS_PASSWORD,
		DB:           0,
		MaxRetries:   3,
		RetryBackoff: 2 * time.Second,
	}
	var err error
	manager.Auth, err = cache.NewCache(authConfig)
	if err != nil {
		log.Fatal(err)
	}
	manager.Api, err = cache.NewCache(apiConfig)
	if err != nil {
		log.Fatal(err)
	}
	manager.Verify, err = cache.NewCache(verifyConfig)
	if err != nil {
		log.Fatal(err)
	}
}
