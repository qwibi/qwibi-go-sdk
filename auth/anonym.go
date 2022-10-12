package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
)

// QAnonymousAuth ...
type QAnonymousAuth struct{}

// NewAnonymousAuth ...
func NewAnonymousAuth() (*QAnonymousAuth, error) {
	auth := &QAnonymousAuth{}
	return auth, nil
}

// NewAnonymousAuthPb ...
func NewAnonymousAuthPb(pb *proto.QPBxAnonymAuth) (*QAnonymousAuth, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	auth := &QAnonymousAuth{}

	return auth, nil
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
