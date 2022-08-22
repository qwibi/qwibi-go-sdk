package response

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/auth"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QAuthResponse ...
type QAuthResponse struct {
	Session *auth.QSession
}

// NewAuthResponse ...
func NewAuthResponse(session *auth.QSession) (*QAuthResponse, error) {

	r := &QAuthResponse{
		Session: session,
	}

	if err := r.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

// NewAuthResponsePb ...
func NewAuthResponsePb(pb *proto.QPBxAuthResponse) (*QAuthResponse, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	res := &QAuthResponse{}

	if err := res.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

// Valid ...
func (c *QAuthResponse) Valid() error {
	return c.Session.Valid()
}

// Pb ...
func (c *QAuthResponse) Pb() (*proto.QPBxAuthResponse, error) {

	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	sessionPB, err := c.Session.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxAuthResponse{
		Session: sessionPB,
	}

	return pb, nil
}
