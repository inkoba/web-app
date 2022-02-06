package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

// Files properties to work with files
type Files struct {
	ErrorLogFile string `mapstructure:"error_log_file"`
	InfoLogFile  string `mapstructure:"info_log_file"`
}

type Mongo struct {
	URI string `mapstructure:"uri"`
}

type Config struct {
	Port    int    `mapstructure:"port"`
	LogMode string `mapstructure:"log_mode"`
	Files   `mapstructure:"files"`
	Mongo   `mapstructure:"mongo"`
}

var instance *singleton = nil
var once sync.Once

type singleton struct {
	conf Config
	sync.RWMutex
}

type Singleton interface {
	Get() Config
	Set(key string, value interface{})
}

func GetInstance() Singleton {
	once.Do(func() {
		instance = new(singleton)
		instance.conf.getConfig()
	})
	return instance
}

func (s *singleton) Get() Config {
	s.RLock()
	defer s.RUnlock()
	return s.conf
}

func (s *singleton) Set(key string, value interface{}) {
	s.Lock()
	defer s.Unlock()

	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		log.Println("Unable to write config ", err)
		return
	}

	if err := viper.Unmarshal(&s.conf); err != nil {
		log.Fatalf("error unmarshall configs: %s", err.Error())
	}

}

func (c *Config) getConfig() *Config {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := viper.Unmarshal(c); err != nil {
		log.Fatalf("error unmarshall configs: %s", err.Error())
	}
	return c

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
