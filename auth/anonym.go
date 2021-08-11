package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

//type AuthMethod int

// QAnonymousAuth ...
type QAnonymousAuth struct {
	//Object *geo.QGeoObject
}

func (c *QAnonymousAuth) Valid() (*QAnonymousAuth, error) {
	return c, nil
}

// Pb ...
func (c *QAnonymousAuth) Pb() (*proto.QPBxAnonymAuth, error) {
	pb := &proto.QPBxAnonymAuth{}
	return pb, nil
}

// NewAnonymousAuth ...
func NewAnonymousAuth() (*QAnonymousAuth, error) {
	auth := &QAnonymousAuth{}
	return auth, nil
}

// NewAnonymousAuthPb ...
func NewAnonymousAuthPb(pb *proto.QPBxAuthRequest_Anonym) (*QAnonymousAuth, error) {
	if pb == nil {
		return nil, errors.New("Invalid parameter type nil")
	}

	auth := &QAnonymousAuth{}
	return auth, nil
}
