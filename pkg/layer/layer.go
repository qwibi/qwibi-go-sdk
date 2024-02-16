package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeoLayer ...
type QGeoLayer struct {
	LayerId    string            `json:"layer_id"`
	Public     bool              `json:"public"`
	Properties []byte            `json:"properties"`
	Commands   map[string]string `json:"commands"`
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
		WithLayerCommands(in.Commands),
	)

	return layer, nil
}

func (c *QGeoLayer) Pb() *proto.QPBxGeoLayer {
	return &proto.QPBxGeoLayer{
		LayerId:    c.LayerId,
		Properties: c.Properties,
		Commands:   c.Commands,
	}
}

func (c *QGeoLayer) Valid() error {
	if c.LayerId == "" {
		return qlog.Error("layer validation error")
	}

	return nil
}
