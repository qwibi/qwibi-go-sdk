package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QJoinResponse ...
type QJoinResponse struct {
	Object geo.QGeoObject
}

// NewJoinResponse ...
func NewJoinResponse(object geo.QGeoObject) (*QJoinResponse, error) {
	if object == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	r := &QJoinResponse{
		Object: object,
	}

	return r, r.Valid()
}

// NewJoinResponsePb ...
func NewJoinResponsePb(pb *proto.QPBxJoinResponse) (*QJoinResponse, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	object, err := geo.NewGeoObjectPb(pb.Object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	r := &QJoinResponse{
		Object: object,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// Valid ...
func (c *QJoinResponse) Valid() error {
	if c.Object == nil {
		err := errors.New("Invalid format")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	if err := c.Object.Valid(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Pb ...
func (c *QJoinResponse) Pb() *proto.QPBxJoinResponse {
	return &proto.QPBxJoinResponse{
		Object: c.Object.Pb(),
	}
}
