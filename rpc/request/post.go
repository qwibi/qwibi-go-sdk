package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QPostRequest ...
type QPostRequest struct {
	Object *geo.QGeoObject
}

// NewPostRequest ...
func NewPostRequest(object *geo.QGeoObject) (*QPostRequest, error) {
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
		err := errors.New("Invalid message type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	if pb.Object == nil {
		err := errors.New("Invalid object type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	req := &QPostRequest{}

	object, err := geo.NewGeoObjectPb(pb.Object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Object = object

	return req, nil
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
func (c *QPostRequest) Pb() (*proto.QPBxPostRequest, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pbObject, err := c.Object.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxPostRequest{
		Object: pbObject,
	}

	return pb, nil
}
