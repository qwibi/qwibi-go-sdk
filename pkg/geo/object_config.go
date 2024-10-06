package geo

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
)

type ObjectOption func(c *QGeoObject) error

func WithGid(gid string) ObjectOption {
	return func(c *QGeoObject) error {
		if gid != "" {
			c.Gid = gid
		} else {
			c.Gid = utils.NewID()
		}
		return nil
	}
}

func WithProperties(b []byte) ObjectOption {
	return func(c *QGeoObject) error {
		if b != nil {
			c.Properties = b
		}
		return nil
	}
}
