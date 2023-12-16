package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

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

func WithLayerCommands(commands ...*command.QCommand) LayerOption {
	return func(c *QGeoLayer) {
		c.commands = commands
	}
}

func WithLayerCommandsPb(commands []*proto.QPBxCommand) LayerOption {
	return func(c *QGeoLayer) {
		cmdList := []*command.QCommand{}
		for _, cmdPb := range commands {
			cmd, err := command.NewCommandPb(cmdPb)
			if err != nil {
				qlog.Error(err)
			}
			cmdList = append(cmdList, cmd)
		}

		c.commands = cmdList
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
