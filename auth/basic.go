package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
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
		return nil, qlog.Error(err)
	}

	return auth, nil
}

// NewBasicAuthPb ...
func NewBasicAuthPb(pb *proto.QPBxBasicAuth) (*QBasicAuth, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	login := pb.Login
	password := pb.Password

	auth, err := NewBasicAuth(login, password)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return auth, nil
}

// Valid ...
func (c *QBasicAuth) Valid() error {
	if c.Login == "" || c.Password == "" {
		err := errors.New("Login or password not defined")
		return qlog.Error(err)
	}

	return nil
}

// Pb ...
func (c *QBasicAuth) Pb() (*proto.QPBxBasicAuth, error) {
	if err := c.Valid(); err != nil {
		return nil, qlog.Error(err)
	}

	pb := &proto.QPBxBasicAuth{
		Login:    c.Login,
		Password: c.Password,
	}

	return pb, nil
}
