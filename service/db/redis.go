package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

type RedisService interface {
	Ping() error
	Set(string, string, time.Time) error
	Get(string) (string, error)
	Del(string) error
	GetExpireAt(string) (time.Time, error)
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

// -------------------------------------------------
func (rs *RedisServiceImpl) Set(key string, token string, duration time.Duration) error {
	err := rs.client.Set(key, token, duration).Err()
	if err != nil {
		rs.Logger.Info("Redis set", "err", err)
		return err
	}
	return nil
}

func (rs *RedisServiceImpl) Get(key string) (string, error) {
	token, err := rs.client.Get(key).Result()
	if err != nil {
		rs.Logger.Info("Redis get", "err", err)
		//if err == redis.ErrNil {
		//	return "", nil
		//}
		return "", err
	}
	return token, nil
}

func (rs *RedisServiceImpl) Del(key string) error {
	err := rs.client.Del(key).Err()
	if err != nil {
		rs.Logger.Info("Redis del", "err", err)
		return err
	}
	return nil
}

func (rs *RedisServiceImpl) GetExpireAt(key string) (time.Time, error) {
	res, err := rs.client.TTL(key).Result()
	if err != nil {
		return time.Time{}, err
	}
	//if res == -2 {
	//}
	//if res == -1 {
	//}
	// to..do.. +4 seconds ttl
	expireAt, err := time.ParseDuration(fmt.Sprintf("%ds", res-4))
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(expireAt), nil
}
