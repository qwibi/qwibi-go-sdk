package auth

import (
	"errors"

	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QAuth interface {
	Valid() error
}

func NewAuthRequestPb(pb *proto.QPBxAuthRequest) (QAuth, error) {
	switch v := pb.Auth.(type) {
	case *proto.QPBxAuthRequest_Anonym:
		return NewAnonymousAuthPb(v.Anonym)
	case *proto.QPBxAuthRequest_Basic:
		return NewBasicAuthPb(v.Basic)
	default:
		return nil, errors.New("Unknown auth request type")
	}
}
