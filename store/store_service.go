package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// wrapper of redis client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cache = 6 * time.Hour

// connect to redis and return a client
func InitilizeStorage() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error init Redis : %v", err)
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}
// save short url
func SaveUrlMapping(shortUrl,originalUrl,userid string) {
	err := storeService.redisClient.Set(ctx,shortUrl,originalUrl,cache)
	
	if err != nil {
		log.Printf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
	}

}
// get the org url
func RetriveInitialUrl (shortUrl string) string {
	result,err := storeService.redisClient.Get(ctx,shortUrl).Result()
	if err != nil {
		log.Printf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl)
	}
	return result
}