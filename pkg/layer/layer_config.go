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

func WithLayerPublic(public bool) LayerOption {
	return func(c *QLayer) {
		c.Public = public
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
