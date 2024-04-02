package redis_test

import (
	"log"
	"os"
	"testing"

	"github.com/0x30c4/Go-Backend-BoilerPlate/internal/env"
	"github.com/0x30c4/Go-Backend-BoilerPlate/internal/logger"
	"github.com/0x30c4/Go-Backend-BoilerPlate/internal/redis"
	"github.com/alicebob/miniredis/v2"
)

func TestRedisConnection(t *testing.T) {
  env.EnvInit()

  // changing the defualt logs log file and
  // setting up the logger
  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  // creating mock redis server
  redisTest := miniredis.RunT(t)

  env.REDIS_ADDRESS = redisTest.Addr()

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  // check if invalid server throws any error or not
  env.REDIS_ADDRESS = "invalid_address:1212"

  err = redis.InitRedis()

  if err == nil {
    t.Errorf("redis.InitRedis() didn't throw any errors")
  }

}
