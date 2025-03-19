package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
)

type LayerOption func(config *QLayer)

func WithGid(gid string) LayerOption {
	return func(c *QLayer) {
		if gid != "" {
			c.Gid = gid
		} else {
			c.Gid = utils.NewID()
		}
	}
}

func WithLaeyrId(layerId string) LayerOption {
	return func(c *QLayer) {
		if layerId != "" {
			c.LayerId = layerId
		} else {
			c.LayerId = utils.NewID()
		}
	}
}

func WithIsPublic(f bool) LayerOption {
	return func(c *QLayer) {
		c.IsPublic = f
	}
}

func WithLayerName(name string) LayerOption {
	return func(c *QLayer) {
		c.Name = name
	}
}

func WithLayerProperties(b []byte) LayerOption {
	return func(c *QLayer) {
		c.Properties = b
	}
}

func WithLayerCommands(commands map[string]string) LayerOption {
	return func(c *QLayer) {
		c.Commands = commands
	}
}
