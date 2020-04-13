package models

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)


func create(file string) *App {
	config := InitServer(file)
	return &App{
		Host: viper.GetString("server.host"),
		Port: viper.GetString("server.port"),
		Server: InitServer(config),
	}
}

type App struct {
	Host, Port string
	Server *http.Server
}

func InitServer(c *Config) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%s", c.Host, c.Port),
		Handler:           c.Handler,
		ReadTimeout:       c.ReadTimeout,
		WriteTimeout:      c.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    c.MaxHeaderBytes,
		ErrorLog:          c.Log,
	}
}

func(a App) run() {
	a.Server.ListenAndServe()
}