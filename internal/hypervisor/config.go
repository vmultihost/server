package hypervisor

import "time"

type Config struct {
	socketName    string
	socketTimeout time.Duration
	imgPath       string
}

func NewConfig(socketName string, socketTimeout time.Duration, imgPath string) *Config {
	return &Config{socketName, socketTimeout, imgPath}
}
