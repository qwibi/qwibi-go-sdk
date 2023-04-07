package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
	"github.com/qwibi/qwibi-go-sdk/utils"
)

// QSession ...
type QSession struct {
	Token string
}

// NewSession ...
func NewSession() (*QSession, error) {
	session := &QSession{
		Token: utils.NewToken(),
	}

	if err := session.Valid(); err != nil {
		return nil, qlog.Error(err)
	}

	return session, nil
}

// NewSessionPb ...
func NewSessionPb(pb *proto.QPBxSession) (*QSession, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	session, err := NewSession()
	if err != nil {
		return nil, qlog.Error(err)
	}

	session.Token = pb.Token

	return session, nil
}

// Valid ...
func (c *QSession) Valid() error {
	if c.Token == "" {
		err := errors.New("Sesion token is not defined")
		return qlog.Error(err)
	}

	return nil
}

// Pb ...
func (c *QSession) Pb() (*proto.QPBxSession, error) {
	if err := c.Valid(); err != nil {
		return nil, qlog.Error(err)
	}

	pb := &proto.QPBxSession{
		Token: c.Token,
	}

	return pb, nil
}
