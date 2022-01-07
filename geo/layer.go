package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
)

// QLayer ...
type QLayer struct {
	Gid  string
	Name string
}

func (c *QLayer) Valid() error {
	if c.Gid == "" || c.Name == "" {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return nil
}

// NewLayer ...
func NewLayer() (*QLayer, error) {
	gid := utils.NewID()
	layer := &QLayer{
		Gid:  gid,
		Name: gid,
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
		Gid:  c.Gid,
		Name: c.Name,
	}

	return pb, nil
}

// NewLayerPb ...
func NewLayerPb(pb *proto.QPBxLayer) (*QLayer, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	layer := &QLayer{
		Gid:  pb.Gid,
		Name: pb.Name,
	}

	if err := layer.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return layer, nil
}
