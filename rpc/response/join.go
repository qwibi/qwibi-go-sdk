package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QJoinResponse ...
type QJoinResponse struct {
	Layer *geo.QGeoLayer
}

// NewJoinResponse ...
func NewJoinResponse(layer *geo.QGeoLayer) (*QJoinResponse, error) {
	if layer == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	r := &QJoinResponse{
		Layer: layer,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// NewJoinResponsePb ...
func NewJoinResponsePb(pb *proto.QPBxJoinResponse) (*QJoinResponse, error) {
	if pb == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	layer, err := geo.NewGeoLayerPb(pb.Layer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	r := &QJoinResponse{
		Layer: layer,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// Valid ...
func (c *QJoinResponse) Valid() error {
	if c.Layer == nil {
		err := errors.New("Invalid format")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	if err := c.Layer.Valid(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Pb ...
func (c *QJoinResponse) Pb() (*proto.QPBxJoinResponse, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pbLayer, err := c.Layer.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxJoinResponse{
		Layer: pbLayer,
	}

	return pb, nil
}
