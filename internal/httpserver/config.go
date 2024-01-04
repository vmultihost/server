package httpserver

type Config struct {
	Port     string `toml:"port"`
	Address  string `toml:"address"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		Address:  "localhost",
		LogLevel: "debug",
	}
}
