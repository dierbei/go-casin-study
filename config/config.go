package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	c    *config
)

type config struct {
	Application struct {
		Mode string `yaml:"network"`
		Port int    `yaml:"address"`
	} `yaml:"application"`

	Mysql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
}

func newConfig() *config {
	return &config{}
}

// GetConfig 读取配置 config/config.yaml
func GetConfig() *config {
	once.Do(func() {
		_viper := viper.New()
		_viper.AddConfigPath("config")
		_viper.SetConfigName("config")
		_viper.SetConfigType("yaml")
		if err := _viper.ReadInConfig(); err != nil {
			panic(err)
		}

		c = newConfig()
		if err := _viper.Unmarshal(c); err != nil {
			log.Fatal(err)
		}
	})

	return c
}
