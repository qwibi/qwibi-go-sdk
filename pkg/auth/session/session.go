package session

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QSession ...
type QSession struct {
	AccessToken string
	AccountId   string
}

// NewSession ...
func NewSession(options ...SessionOption) *QSession {
	h := &QSession{
		AccessToken: utils.NewID(),
	}

	for _, opt := range options {
		opt(h)
	}

	return h
}

// NewSessionPb ...
func NewSessionPb(in *proto.QPBxSession) (*QSession, error) {
	if in == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	session := NewSession(
		WithSessionAccessToken(in.AccessToken),
		WithSessionAccountId(in.AccountId),
	)

	return session, nil
}

// Valid ...
func (c *QSession) Valid() error {
	if c.AccessToken == "" {
		return qlog.Error("session token is not defined")
	}

	if c.AccountId == "" {
		return qlog.Error("session account ID is not define")
	}

	return nil
}

// Pb ...
func (c *QSession) Pb() *proto.QPBxSession {
	return &proto.QPBxSession{
		AccessToken: c.AccessToken,
		AccountId:   c.AccountId,
	}
}
