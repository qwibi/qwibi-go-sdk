package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QLayer ...
type QLayer struct {
	LayerID string
	Name    string
}

func (c *QLayer) Valid() error {
	if c.LayerID == "" || c.Name == "" {
		return errors.New("Layer ID or Name is empty")
	}

	return nil
}

// NewLayer ...
func NewLayer(layerID string) (*QLayer, error) {
	layer := &QLayer{
		LayerID: layerID,
		Name:    layerID,
	}

	if err := layer.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return layer, nil
}

// Pb ...
func (c *QLayer) Pb() (*proto.QPBxLayer, error) {

	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxLayer{
		LayerID: c.LayerID,
		Name:    c.Name,
	}

	return pb, nil
}

// NewLayerPb ...
func NewLayerPb(pb *proto.QPBxLayer) (*QLayer, error) {
	if pb == nil {
		return nil, errors.New("Invalid parameter type nil")
	}

	layer := &QLayer{
		LayerID: pb.LayerID,
		Name:    pb.Name,
	}

	if err := layer.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return layer, nil
}
