package apiserver

// Config struct...
type Config struct {
	BindAddr string `toml:"bind_addr"` // адрес на котором запуск сервиса
	LogLevel string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
}

// New Config ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
