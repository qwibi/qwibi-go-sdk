package auth

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"

	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QAuth interface {
	Pb() *proto.QPBxAuth
	//Type() string
	//Valid() error
}

func NewAuthPb(pb *proto.QPBxAuth) (QAuth, error) {
	switch v := pb.Type.(type) {
	case *proto.QPBxAuth_Anonym:
		return NewAnonymousAuthPb(v.Anonym)
	case *proto.QPBxAuth_Basic:
		return NewBasicAuthPb(v.Basic)
	default:
		return nil, qlog.Errorf("unknown auth request type: %+T", v)
	}
}
