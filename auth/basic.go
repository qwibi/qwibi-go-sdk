package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QBasicAuth ...
type QBasicAuth struct {
	Login    string
	Password string
}

func (c *QBasicAuth) Valid() error {
	if c.Login == "" || c.Password == "" {
		return errors.New("Login or password not defined")
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
func NewBasicAuthPb(pb *proto.QPBxAuthRequest_Basic) (*QBasicAuth, error) {
	if pb == nil || pb.Basic == nil {
		return nil, errors.New("Invalid parameter type nil")
	}

	login := pb.Basic.Login
	password := pb.Basic.Password

	auth, err := NewBasicAuth(login, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return auth, nil
}
