package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QPostRequest ...
type QPostRequest struct {
	object *geo.QGeoObject
}

func (c *QPostRequest) Valid() error {
	if c.object == nil {
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	if err := c.object.Valid(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Pb...
func (c *QPostRequest) Pb() (*proto.QPBxPostRequest, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pbObject, err := c.object.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxPostRequest{
		Object: pbObject,
	}

	return pb, nil
}

func NewPostRequest(object *geo.QGeoObject) (*QPostRequest, error) {
	r := &QPostRequest{
		object: object,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}
