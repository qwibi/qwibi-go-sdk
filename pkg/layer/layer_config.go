package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
)

type LayerOption func(config *QGeoLayer)

func WithLayerGid(gid string) LayerOption {
	return func(c *QGeoLayer) {
		if gid != "" {
			c.LayerId = gid
		} else {
			c.LayerId = utils.NewID()
		}

	}
}

func WithLayerCommands(commands map[string]string) LayerOption {
	return func(c *QGeoLayer) {
		c.Commands = commands
	}
}

//func WithLayerCommandsPb(commands map[string]string) LayerOption {
//	return func(c *QGeoLayer) {
//		cmdList := []*command.QCommand{}
//		for _, cmdPb := range commands {
//			cmd, err := command.NewCommandPb(cmdPb)
//			if err != nil {
//				qlog.Error(err)
//			}
//			cmdList = append(cmdList, cmd)
//		}
//
//		c.Commands = comm
//	}
//}

func WithLayerProperties(properties []byte) LayerOption {
	return func(c *QGeoLayer) {
		c.Properties = properties
	}
}

func WithLayerPublic(f bool) LayerOption {
	return func(c *QGeoLayer) {
		c.Public = f
	}
}
