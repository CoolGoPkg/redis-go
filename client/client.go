package client

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	goredis "gopkg.in/redis.v5"
	"time"
)

var LocalRedis *goredis.Client

// 	初始化本地redis
func Init() {
	conf := new(config)
	LoadConf(conf, "../conf/conf.yaml")
	fmt.Println(conf)
	LocalRedis = conf.LocalRedis.NewRedisClient()
}

type RedisConfig struct {
	Host string
	Port int
	Auth string
	DB   int
}

func (c RedisConfig) NewRedisClient() *goredis.Client {
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	cli := NewRedisClient(address, c.Auth, c.DB)
	if err := cli.Ping().Err(); err != nil {
		panic(err)
	}
	return cli
}

func (c RedisConfig) NewRedisPool() *redis.Pool {
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return NewRedisPool(address, c.Auth, c.DB)
}

func NewRedisClient(address, password string, db int) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}

func NewRedisPool(address, password string, db int) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address,
				redis.DialPassword(password),
				redis.DialDatabase(db),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:   500,
		MaxActive: 2000,
		Wait:      true,
	}
}

type RedisClusterConfig struct {
	Addrs       []string `yaml:"addrs"`
	Auth        string   `yaml:"auth"`
	IdleTimeout int      `yaml:"idle_timeout"`
}

func NewRedisCluster(config RedisClusterConfig) *goredis.ClusterClient {
	tradeRedisClusterClient := goredis.NewClusterClient(&goredis.ClusterOptions{
		Addrs:    config.Addrs,
		Password: config.Auth,
		PoolSize: 5000,
	})
	if err := tradeRedisClusterClient.Ping().Err(); err != nil {
		panic(err)
	}
	return tradeRedisClusterClient
}
