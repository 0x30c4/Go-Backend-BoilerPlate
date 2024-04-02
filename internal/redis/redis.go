package redis

import (
	"log"
	"context"
	"github.com/0x30c4/Go-Backend-BoilerPlate/internal/env"
	"github.com/redis/go-redis/v9"
)

type PasteModel struct {
  PasteId   string `redis:"PasteId"`
  BurnAfter uint64 `redis:"BurnAfter"`
  ReadCount uint64 `redis:"ReadCount"`
  DeepUrl   uint8  `redis:"DeepUrl"`
  Secret    string `redis:"Secret"`
}

var REDIS_CLIENT *redis.Client

var ctx = context.Background()

func InitRedis() error {
  // Connect to Redis server
  REDIS_CLIENT = redis.NewClient(&redis.Options{
      Addr:     env.REDIS_ADDRESS, // Redis server address
      Password: "",               // No password set
      DB:       0,                // Use default DB
  })

  // Ping the Redis server to check if connection was successful
  _, err := REDIS_CLIENT.Ping(ctx).Result()

  if err != nil {
    log.Println(err)
    return err
  }
  return nil
}
