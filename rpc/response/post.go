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

func NewPostResponsePb(response *proto.QPBxPostResponse) (*QPostResponse, error) {
	object, err := geo.NewGeoObjectPb(response.Object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	r := &QPostResponse{
		Object: object,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}
