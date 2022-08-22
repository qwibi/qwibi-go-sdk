package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QPostRequest ...
type QPostRequest struct {
	Object geo.QGeoObject
}

// NewPostRequest ...
func NewPostRequest(object geo.QGeoObject) (*QPostRequest, error) {
	r := &QPostRequest{
		Object: object,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// NewPostRequestPb ...
func NewPostRequestPb(pb *proto.QPBxPostRequest) (*QPostRequest, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	if pb.Object == nil {
		err := errors.New("Invalid Object type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	object, err := geo.NewGeoObjectPb(pb.Object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req := &QPostRequest{
		Object: object,
	}

	return req, req.Valid()
}

// Valid ...
func (c *QPostRequest) Valid() error {
	if c.Object == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	if err := c.Object.Valid(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Pb...
func (c *QPostRequest) Pb() *proto.QPBxPostRequest {
	return &proto.QPBxPostRequest{
		Object: c.Object.Pb(),
	}
}
