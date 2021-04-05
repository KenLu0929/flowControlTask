package rdb

import (
	"github.com/go-redis/redis/v8"

	"context"
	"fmt"

	"github.com/KenLu0929/flowControlTask/config"
)

var Ctx = context.Background()
var Rdb *redis.Client

func init() {
	rdbConfig := config.Config.RDB
	addr := fmt.Sprintf("%v:%v", rdbConfig.Host, rdbConfig.Port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: rdbConfig.Password,
		DB:       rdbConfig.Db,
	})

	_, err := Rdb.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}
}
