package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
	"github.com/qwibi/qwibi-go-sdk/utils"
)

// QSession ...
type QSession struct {
	Token   string
	LayerID string
}

// NewSession ...
func NewSession(token string) (*QSession, error) {
	session := &QSession{
		Token:   token,
		LayerID: utils.NewID(),
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

	session, err := NewSession(pb.Token)
	if err != nil {
		return nil, qlog.Error(err)
	}

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
