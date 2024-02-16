package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
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
func NewBasicAuthPb(in *proto.QPBxBasicAuth) (*QBasicAuth, error) {
	if in == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	auth, err := NewBasicAuth(in.Login, in.Password)
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
	if c.Login == "" || c.Password == "" {
		err := errors.New("login or password not defined")
		return qlog.Error(err)
	}

	if err := utils.IsValidLogin(c.Login); err != nil {
		return qlog.Error(err)
	}

	if err := utils.IsValidLogin(c.Password); err != nil {
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
