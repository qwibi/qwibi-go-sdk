package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
)

type LayerOption func(config *QLayer)

func WithLayerGid(gid string) LayerOption {
	return func(c *QLayer) {
		if gid != "" {
			c.LayerId = gid
		} else {
			c.LayerId = utils.NewID()
		}
	}
}

func WithLayerName(name string) LayerOption {
	return func(c *QLayer) {
		c.Name = name
	}
}

func WithLayerPublic(f bool) LayerOption {
	return func(c *QLayer) {
		c.Public = f
	}
}

func WithLayerProperties(properties []byte) LayerOption {
	return func(c *QLayer) {
		c.Properties = properties
	}
}

func WithLayerCommands(commands map[string]string) LayerOption {
	return func(c *QLayer) {
		c.Commands = commands
	}
}
