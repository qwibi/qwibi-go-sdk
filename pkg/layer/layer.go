package layer

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QLayer ...
type QLayer struct {
	Gid        string            `json:"gid"`
	LayerId    string            `json:"layer_id"`
	IsPublic   bool              `json:"is_public"`
	Name       string            `json:"name"`
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

func NewLayerPb(in *proto.QPBxLayer) (*QLayer, error) {
	layer := NewLayer(
		WithGid(in.Gid),
		WithLaeyrId(in.LayerId),
		WithIsPublic(in.IsPublic),
		WithLayerName(in.Name),
		WithLayerProperties(in.Properties),
		WithLayerCommands(in.Commands),
	)

	return layer, nil
}

func (c *QLayer) Pb() *proto.QPBxLayer {
	return &proto.QPBxLayer{
		Gid:        c.Gid,
		LayerId:    c.LayerId,
		IsPublic:   c.IsPublic,
		Name:       c.Name,
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
