package hypervisor

import "time"

type Config struct {
	socketName    string
	socketTimeout time.Duration
}

func NewConfig(socketName string, socketTimeout time.Duration) *Config {
	return &Config{socketName, socketTimeout}
}
