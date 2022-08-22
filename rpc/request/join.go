package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QJoinRequest ...
type QJoinRequest struct {
	Gid string
}

// NewJoinRequest ...
func NewJoinRequest(gid string) (*QJoinRequest, error) {
	if gid == "" {
		err := errors.New("Join Gid not defined")
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}

	r := &QJoinRequest{
		Gid: gid,
	}

	return r, r.Valid()
}

// NewJoinRequestPb ...
func NewJoinRequestPb(pb *proto.QPBxJoinRequest) (*QJoinRequest, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	r, err := NewJoinRequest(pb.Gid)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// Valid ...
func (c *QJoinRequest) Valid() error {
	return nil
}

// Pb ...
func (c *QJoinRequest) Pb() (*proto.QPBxJoinRequest, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxJoinRequest{
		Gid: c.Gid,
	}

	return pb, nil
}
