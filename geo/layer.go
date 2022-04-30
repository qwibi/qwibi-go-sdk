package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

// QGeoLayer ...
type QGeoLayer struct {
	Gid        string
	Properties *structpb.Struct
}

// NewGeoLayer ...
func NewGeoLayer() (*QGeoLayer, error) {
	layer := &QGeoLayer{
		Gid:        utils.NewID(),
		Properties: &structpb.Struct{},
	}

	return layer, nil
}

// NewLayerPb ...
func NewGeoLayerPb(pb *proto.QPBxGeoLayer) (*QGeoLayer, error) {
	if pb == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	layer := &QGeoLayer{
		Gid:        pb.Gid,
		Properties: &structpb.Struct{},
	}

	if err := layer.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return layer, nil
}

// Valid ...
func (c *QGeoLayer) Valid() error {
	if c.Gid == "" {
		err := errors.New("Layer gid is not defined")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return nil
}

// Pb ...
func (c *QGeoLayer) Pb() (*proto.QPBxGeoLayer, error) {

	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxGeoLayer{
		Gid:        c.Gid,
		Properties: c.Properties,
	}

	return pb, nil
}
