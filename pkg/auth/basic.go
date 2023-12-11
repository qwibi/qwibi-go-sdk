package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QBasicAuth ...
type QBasicAuth struct {
	AccountId string
	Password  string
}

// NewBasicAuth ...
func NewBasicAuth(login string, password string) (*QBasicAuth, error) {
	auth := &QBasicAuth{
		AccountId: login,
		Password:  password,
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

	login := pb.AccountId
	password := pb.Password

	auth, err := NewBasicAuth(login, password)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return auth, nil
}

// Type ...
func (c *QBasicAuth) Type() string {
	return QBasicAuthType
}

// Valid ...
func (c *QBasicAuth) Valid() error {
	if c.AccountId == "" || c.Password == "" {
		err := errors.New("AccountId or password not defined")
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
		AccountId: c.AccountId,
		Password:  c.Password,
	}

	return pb, nil
}
