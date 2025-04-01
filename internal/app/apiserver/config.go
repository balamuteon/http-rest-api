package apiserver

import "github.com/balamuteon/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` // адрес на котором запуск сервиса
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// New Config ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
