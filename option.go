package config

type OptFunc func(*Config)

func WithEnvPath(path string) OptFunc {
	return func(c *Config) {
		c.envPath = path
	}
}

func WithPrefix(prefix string) OptFunc {
	return func(c *Config) {
		c.prefix = prefix
	}
}
