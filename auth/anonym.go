package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QAnonymousAuth ...
type QAnonymousAuth struct {
}

// NewAnonymousAuth ...
func NewAnonymousAuth() (*QAnonymousAuth, error) {
	auth := &QAnonymousAuth{}
	return auth, nil
}

// NewAnonymousAuthPb ...
func NewAnonymousAuthPb(pb *proto.QPBxAuthRequest_Anonym) (*QAnonymousAuth, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	auth := &QAnonymousAuth{}
	return auth, nil
}

// Valid ...
func (c *QAnonymousAuth) Valid() (*QAnonymousAuth, error) {
	return c, nil
}

// Pb ...
func (c *QAnonymousAuth) Pb() (*proto.QPBxAnonymAuth, error) {
	pb := &proto.QPBxAnonymAuth{}
	return pb, nil
}
