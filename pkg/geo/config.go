package geo

type Config struct {
	Gid        string
	Properties []byte
}

type Option func(config *Config)

func WithGid(gid string) Option {
	return func(c *Config) {
		c.Gid = gid
	}
}

func WithProperties(properties []byte) Option {
	return func(c *Config) {
		c.Properties = properties
	}
}
