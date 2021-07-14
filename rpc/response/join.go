package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QJoinResponse ...
type QJoinResponse struct {
	Layer *geo.QLayer
}

func (c *QJoinResponse) Valid() error {
	if c.Layer == nil {
		return errors.New("Layer is not defined")
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

	pbObject, err := c.Layer.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxJoinResponse{
		Layer: pbObject,
	}

	return pb, nil
}

// NewJoinResponse ...
func NewJoinResponse(layer *geo.QLayer) (*QJoinResponse, error) {
	if layer == nil {
		return nil, errors.New("Wrong layer parameter type nil")
	}

	r := &QJoinResponse{
		Layer: layer,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}
