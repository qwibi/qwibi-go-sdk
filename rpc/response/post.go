package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPostResponse ...
type QPostResponse struct {
	Object *geo.QGeoObject
}

// NewPostResponse ...
func NewPostResponse(object *geo.QGeoObject) (*QPostResponse, error) {
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
func (c *QPostResponse) Pb() (*proto.QPBxPostResponse, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pbObject, err := c.Object.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxPostResponse{
		Object: pbObject,
	}

	return pb, nil
}
