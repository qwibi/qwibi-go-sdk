package geo

type Config struct {
	Gid string
}

type Option func(config *Config)

func WithGid(gid string) Option {
	return func(c *Config) {
		c.Gid = gid
	}
}
