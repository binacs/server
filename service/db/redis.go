package db

import (
	"github.com/go-redis/redis"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

type RedisService interface {
}

type RedisServiceImpl struct {
	Config  *config.Config `inject-name:"Config"`
	Logger  log.Logger     `inject-name:"RedisLogger"`
	client  *redis.Client
	fClient *redis.Client
	name    string
	// state
}

func (rs *RedisServiceImpl) AfterInject() error {
	var err error
	rs.client, err = NewRedisCli(rs.Config.RedisConfig)
	if err != nil {
		return err
	}
	rs.fClient, err = NewRedisSentinelCli(rs.Config.RedisConfig)
	if err != nil {
		return err
	}
	return nil
}

func NewRedisCli(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Network:      cfg.Network,
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	//if _, err := client.Ping().Result(); err != nil {
	//	return nil, fmt.Errorf("NewRedisCli err: %v", err)
	//}
	return client, nil
}

func NewRedisSentinelCli(cfg config.RedisConfig) (*redis.Client, error) {
	failoverOpt := &redis.FailoverOptions{
		MasterName:    cfg.SentinelMaster,
		SentinelAddrs: cfg.SentinelEndpoints,
		Password:      cfg.Password,
		DB:            cfg.DB,
		MaxRetries:    2,
		PoolSize:      cfg.PoolSize,
		OnConnect: func(cn *redis.Conn) error {
			return nil
		},
	}
	client := redis.NewFailoverClient(failoverOpt)
	//if _, err := client.Ping().Result(); err != nil {
	//	return nil, fmt.Errorf("NewRedisSentinelCli err: %v", err)
	//}
	return client, nil
}
