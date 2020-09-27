package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Port int //handle port
}

func defaultConfig() *Config {
	return &Config{
		Port: 18080,
	}
}

type Options func(config *Config)

func ParseConfig(opts []Options) *Config {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func Store(path string, config *Config) error {
	marshal, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, marshal, 0755)
}

func Load(path string, config *Config) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, config)
}
