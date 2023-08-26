package object

import "github.com/qwibi/qwibi-go-sdk/pkg/geometry"

type Config struct {
	Gid        string
	Geometry   *geometry.QGeometry
	Properties []byte
}

type Option func(config *Config)

func WithGid(gid string) Option {
	return func(c *Config) {
		c.Gid = gid
	}
}

func WithGeometry(geometry *geometry.QGeometry) Option {
	return func(c *Config) {
		c.Geometry = geometry
	}
}

func WithProperties(properties []byte) Option {
	return func(c *Config) {
		c.Properties = properties
	}
}
