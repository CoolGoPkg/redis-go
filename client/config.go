package client

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type config struct {
	LocalRedis RedisConfig `yaml:"local_redis"`
}

// LoadConf loads configuration from specified file
func LoadConf(dest interface{}, path string) {
	if err := configor.Load(dest, path); err != nil {
		panic(fmt.Sprintf("failed to load local config file: %v", err))
	}
}
