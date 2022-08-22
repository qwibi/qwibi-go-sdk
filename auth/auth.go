package auth

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

type Auth interface {
	Valid() error
}

func NewAuthPb(pb *proto.QPBxAuthRequest) (Auth, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	switch v := pb.Auth.(type) {
	case *proto.QPBxAuthRequest_Anonym:
		return NewAnonymousAuthPb(v.Anonym)
	case *proto.QPBxAuthRequest_Basic:
		return NewBasicAuthPb(v.Basic)
	default:
		return nil, errors.Errorf("Unknown auth method type %T", v)
	}
}
