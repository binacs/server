package service

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/binacsgo/log"

	"github.com/binacs/server/config"
)

func newRedisCli(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Network:      cfg.Network,
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	return client, nil
}

func newRedisSentinelCli(cfg config.RedisConfig) (*redis.Client, error) {
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
	return client, nil
}

// ----------------------------------------------------------------------

// RedisServiceImpl inplement of RedisService
type RedisServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"RedisLogger"`
	client *redis.Client
	name   string
	// state
}

// AfterInject do inject
func (rs *RedisServiceImpl) AfterInject() error {
	// Ignore the Ping error in AfterInject
	_ = rs.buildClient()

	go rs.checkLoop()
	return nil
}

func (rs *RedisServiceImpl) buildClient() (err error) {
	if rs.Config.RedisConfig.SentinelOn {
		rs.client, err = newRedisSentinelCli(rs.Config.RedisConfig)
	} else {
		rs.client, err = newRedisCli(rs.Config.RedisConfig)
	}
	return err
}

func (rs *RedisServiceImpl) checkLoop() {
	timer := time.NewTimer(dbCheckInterval)
	defer timer.Stop()
	for {
		timer.Reset(dbCheckInterval)
		select {
		case <-timer.C:
			if err := rs.Ping(); err != nil {
				rs.Logger.Error("RedisServiceImpl checkLoop", "err", err)
				if err = rs.buildClient(); err != nil {
					rs.Logger.Error("RedisServiceImpl checkLoop", "err", err)
				}
			}
		}
	}
}

// Ping check the connection
func (rs *RedisServiceImpl) Ping() error {
	if msg, err := rs.client.Ping().Result(); err != nil {
		return fmt.Errorf("RedisCli Ping: msg: %s, err: %v", msg, err)
	}
	return nil
}

// Set set key
func (rs *RedisServiceImpl) Set(key string, token string, duration time.Duration) error {
	err := rs.client.Set(key, token, duration).Err()
	if err != nil {
		rs.Logger.Info("Redis set", "err", err)
		return err
	}
	return nil
}

// Get get key
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

// Del delete key
func (rs *RedisServiceImpl) Del(key string) error {
	err := rs.client.Del(key).Err()
	if err != nil {
		rs.Logger.Info("Redis del", "err", err)
		return err
	}
	return nil
}

// GetExpireAt get key expire-time
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
