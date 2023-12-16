package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeoLayer ...
type QGeoLayer struct {
	layerId    string              `json:"layerId"`
	public     bool                `json:"public"`
	properties []byte              `json:"properties"`
	commands   []*command.QCommand `json:"commands"`
}

// NewGeoLayer ...
func NewGeoLayer(options ...LayerOption) *QGeoLayer {
	h := &QGeoLayer{}

	for _, opt := range options {
		opt(h)
	}

	return h
}

func NewGeoLayerPb(in *proto.QPBxGeoLayer) (*QGeoLayer, error) {
	layer := NewGeoLayer(
		WithLayerGid(in.LayerId),
		WithLayerProperties(in.Properties),
		WithLayerCommandsPb(in.Commands),
	)
	return layer, nil

}

func (c *QGeoLayer) LayerId() string {
	return c.layerId
}

func (c *QGeoLayer) Public() bool {
	return c.public
}

func (c *QGeoLayer) Properties() []byte {
	return c.properties
}

func (c *QGeoLayer) Commands() []*command.QCommand {
	return c.commands
}

func (c *QGeoLayer) Pb() *proto.QPBxGeoLayer {
	commandsPb := []*proto.QPBxCommand{}
	for _, cmdPb := range c.commands {
		commandsPb = append(commandsPb, cmdPb.Pb())
	}

	return &proto.QPBxGeoLayer{
		LayerId:    c.layerId,
		Properties: c.properties,
		Commands:   commandsPb,
	}
}

func (c *QGeoLayer) Valid() error {
	if c.layerId == "" {
		return qlog.Error("layer validation error")
	}

	return nil
}
