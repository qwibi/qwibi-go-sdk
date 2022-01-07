package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QJoinRequest ...
type QJoinRequest struct {
	LayerID string
}

func (c *QJoinRequest) Valid() error {
	if c.LayerID == "" {
		err := errors.New("Invalid format")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
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
		err := errors.New("Invalid format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	r := &QJoinRequest{
		LayerID: pb.LayerID,
	}

	return r, nil
}
