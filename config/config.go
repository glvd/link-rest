package config

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
