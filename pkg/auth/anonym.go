package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QAnonymousAuth ...
type QAnonymousAuth struct {
	Token string
}

// NewAnonymousAuth ...
func NewAnonymousAuth(options ...AnonymousAuthOption) (*QAnonymousAuth, error) {
	h := &QAnonymousAuth{}

	for _, opt := range options {
		opt(h)
	}

	return h, nil
}

// NewAnonymousAuthPb ...
func NewAnonymousAuthPb(in *proto.QPBxAnonymAuth) (*QAnonymousAuth, error) {
	if in == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	auth, err := NewAnonymousAuth()
	if err != nil {
		return nil, qlog.Error(err)
	}

	return auth, nil
}

// Type ...
func (c *QAnonymousAuth) Type() string {
	return QAnonymousAuthType
}

// Valid ...
func (c *QAnonymousAuth) Valid() error {
	return nil
}

// Pb ...
func (c *QAnonymousAuth) Pb() (*proto.QPBxAnonymAuth, error) {
	pb := &proto.QPBxAnonymAuth{}
	return pb, nil
}
