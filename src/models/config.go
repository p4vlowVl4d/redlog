package models

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type (
	Config struct {
		Host, Port     string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		Handler		   http.Handler
		MaxHeaderBytes int
		Log            *log.Logger
	}
)
func(c Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func InitConfig(config string) *Config {
	if (! FileExists(config)) {
		panic("Файл не найден")
	}
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	return &Config{
		Host:           "",
		Port:           "",
		ReadTimeout:    0,
		WriteTimeout:   0,
		Handler:        nil,
		MaxHeaderBytes: 0,
		Log:            nil,
	}
}
