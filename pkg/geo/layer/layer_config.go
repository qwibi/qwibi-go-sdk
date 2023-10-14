package layer

import "github.com/qwibi/qwibi-go-sdk/pkg/utils"

type LayerOption func(config *QGeoLayer)

func WithLayerGid(gid string) LayerOption {
	return func(c *QGeoLayer) {
		if gid != "" {
			c.layerId = gid
		} else {
			c.layerId = utils.NewID()
		}

	}
}

func WithLayerProperties(properties []byte) LayerOption {
	return func(c *QGeoLayer) {
		c.properties = properties
	}
}

func WithLayerPublic(f bool) LayerOption {
	return func(c *QGeoLayer) {
		c.public = f
	}
}
