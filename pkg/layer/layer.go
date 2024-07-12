package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QLayer ...
type QLayer struct {
	LayerId    string            `json:"layer_id"`
	Name       string            `json:"name"`
	Public     bool              `json:"public"`
	Properties []byte            `json:"properties"`
	Commands   map[string]string `json:"commands"`
}

// NewLayer ...
func NewLayer(options ...LayerOption) *QLayer {
	h := &QLayer{}

	for _, opt := range options {
		opt(h)
	}

	return h
}

func NewGeoLayerPb(in *proto.QPBxLayer) (*QLayer, error) {
	layer := NewLayer(
		WithLayerGid(in.LayerId),
		WithLayerName(in.Name),
		WithLayerPublic(in.Public),
		WithLayerProperties(in.Properties),
		WithLayerCommands(in.Commands),
	)

	return layer, nil
}

func (c *QLayer) Pb() *proto.QPBxLayer {
	return &proto.QPBxLayer{
		LayerId:    c.LayerId,
		Name:       c.Name,
		Public:     c.Public,
		Properties: c.Properties,
		Commands:   c.Commands,
	}
}

func (c *QLayer) Valid() error {
	if c.LayerId == "" {
		return qlog.Error("layer id is not defined")
	}

	if c.Name == "" {
		return qlog.Error("layer name is not defined")
	}

	return nil
}
