package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QJoinRequest ...
type QJoinRequest struct {
	LayerID string
}

func (c *QJoinRequest) Valid() error {
	if c.LayerID == "" {
		return errors.New("LayerID is not defined")
	}

	return nil
}

// Pb ...
func (c *QJoinRequest) Pb() (*proto.QPBxJoinRequest, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxJoinRequest{
		LayerID: c.LayerID,
	}

	return pb, nil
}

// NewJoinRequest ...
func NewJoinRequest(layerID string) (*QJoinRequest, error) {
	r := &QJoinRequest{
		LayerID: layerID,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// NewJoinRequestPb ...
func NewJoinRequestPb(pb *proto.QPBxJoinRequest) (*QJoinRequest, error) {
	if pb == nil {
		return nil, errors.New("Invalid format")
	}

	r := &QJoinRequest{
		LayerID: pb.LayerID,
	}

	return r, nil
}
