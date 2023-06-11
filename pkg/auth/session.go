package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/internal/utils"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QSession ...
type QSession struct {
	Token string
	//Gid   string
}

// NewSession ...
func NewSession() (*QSession, error) {
	session := &QSession{
		Token: utils.NewToken(),
		//Gid:   utils.NewID(),
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
		return qlog.Error("Session token is not defined")
	}

	//if c.Gid == "" {
	//	return qlog.Error("Session gid is not defined")
	//}

	return nil
}

// Pb ...
func (c *QSession) Pb() (*proto.QPBxSession, error) {
	if err := c.Valid(); err != nil {
		return nil, qlog.Error(err)
	}

	pb := &proto.QPBxSession{
		Token: c.Token,
		//Gid:   c.Gid,
	}

	return pb, nil
}