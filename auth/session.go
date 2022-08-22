package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QSession ...
type QSession struct {
	Token string
}

// NewSession ...
func NewSession(token string) (*QSession, error) {
	session := &QSession{
		Token: token,
	}

	if err := session.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

// NewSessionPb ...
func NewSessionPb(pb *proto.QPBxSession) (*QSession, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	session, err := NewSession(pb.Token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

// Valid ...
func (c *QSession) Valid() error {
	if c.Token == "" {
		err := errors.New("Sesion token is not defined")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return nil
}

// Pb ...
func (c *QSession) Pb() (*proto.QPBxSession, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxSession{
		Token: c.Token,
	}

	return pb, nil
}
