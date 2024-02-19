package httpserver

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	port     string
	address  string
	logLevel string
}

type configDto struct {
	Port     string `toml:"port"`
	Address  string `toml:"address"`
	LogLevel string `toml:"log_level"`
}

func NewDefaultConfig() *Config {
	return &Config{
		address:  "localhost",
		port:     "8080",
		logLevel: "debug",
	}
}

func NewConfig(address, port, logLevel string) *Config {
	return &Config{
		address:  address,
		port:     port,
		logLevel: logLevel,
	}
}

func NewConfigFromFile(configPath string) (*Config, error) {
	config := NewDefaultConfig().toDto()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		return nil, err
	}

	return config.fromDto(), nil
}

func (c *configDto) fromDto() *Config {
	return NewConfig(c.Address, c.Port, c.LogLevel)
}

func (c *Config) toDto() *configDto {
	return &configDto{
		Address:  c.address,
		Port:     c.port,
		LogLevel: c.logLevel,
	}
}

func (c *Config) FullAddress() string {
	return fmt.Sprintf("%s:%s", c.address, c.port)
}
