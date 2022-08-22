package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QPostResponse ...
type QPostResponse struct {
	Object geo.QGeoObject
}

// NewPostResponse ...
func NewPostResponse(object geo.QGeoObject) (*QPostResponse, error) {
	if object == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	res := &QPostResponse{
		Object: object,
	}

	if err := res.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

// NewPostResponsePb ...
func NewPostResponsePb(pb *proto.QPBxPostResponse) (*QPostResponse, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	object, err := geo.NewGeoObjectPb(pb.Object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := &QPostResponse{
		Object: object,
	}

	if err := res.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

// Valid ...
func (c *QPostResponse) Valid() error {
	if err := c.Object.Valid(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Pb...
func (c *QPostResponse) Pb() *proto.QPBxPostResponse {
	return &proto.QPBxPostResponse{
		Object: c.Object.Pb(),
	}
}
