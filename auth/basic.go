package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QBasicAuth ...
type QBasicAuth struct {
	Login    string
	Password string
}

// NewBasicAuth ...
func NewBasicAuth(login string, password string) (*QBasicAuth, error) {
	auth := &QBasicAuth{
		Login:    login,
		Password: password,
	}

	if err := auth.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	return auth, nil
}

// NewBasicAuthPb ...
func NewBasicAuthPb(pb *proto.QPBxBasicAuth) (*QBasicAuth, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	login := pb.Login
	password := pb.Password

	auth, err := NewBasicAuth(login, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return auth, nil
}

// Valid ...
func (c *QBasicAuth) Valid() error {
	if c.Login == "" || c.Password == "" {
		err := errors.New("Login or password not defined")
		log.Error().Stack().Err(err).Msg("")
		return errors.WithStack(err)
	}

	return nil
}

// Pb ...
func (c *QBasicAuth) Pb() (*proto.QPBxBasicAuth, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxBasicAuth{
		Login:    c.Login,
		Password: c.Password,
	}

	return pb, nil
}
