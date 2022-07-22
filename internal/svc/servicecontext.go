package svc

import (
	"awesome/internal/config"
	"github.com/dtm-labs/rockscache"
	"github.com/go-redis/redis/v8"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	Redis      *redis.Client
	RocksCache *rockscache.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	rockscache.SetVerbose(true)
	return &ServiceContext{
		Config: c,
		Redis:  redisClient,
		RocksCache: rockscache.NewClient(redisClient, rockscache.Options{
			StrongConsistency:      true,
			Delay:                  10 * time.Second,
			EmptyExpire:            60 * time.Second,
			LockExpire:             3 * time.Second,
			LockSleep:              100 * time.Millisecond,
			RandomExpireAdjustment: 0.1,
			WaitReplicasTimeout:    3000 * time.Millisecond,
		}),
	}
}
