package request

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/auth"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
)

// QAuthRequest ...
type QAuthRequest struct {
	Auth interface{}
}

// NewAuthRequest ...
func NewAuthRequest(authMethod interface{}) (*QAuthRequest, error) {
	switch v := authMethod.(type) {
	case *auth.QAnonymousAuth:
		req := &QAuthRequest{
			Auth: v,
		}
		return req, nil
	case *auth.QBasicAuth:
		req := &QAuthRequest{}
		return req, nil
	default:
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}

// NewAuthRequestPb ...
func NewAuthRequestPb(pb *proto.QPBxAuthRequest) (*QAuthRequest, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	switch v := pb.Auth.(type) {
	case *proto.QPBxAuthRequest_Anonym:
		t, err := auth.NewAnonymousAuthPb(v.Anonym)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &QAuthRequest{Auth: t}, nil
	case *proto.QPBxAuthRequest_Basic:
		t, err := auth.NewBasicAuthPb(v.Basic)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &QAuthRequest{Auth: t}, nil
	default:
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}

// Pb ...
func (c *QAuthRequest) Pb() (*proto.QPBxAuthRequest, error) {
	switch v := c.Auth.(type) {
	case *auth.QAnonymousAuth:
		pbAuth, err := v.Pb()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		request := &proto.QPBxAuthRequest{
			Auth: &proto.QPBxAuthRequest_Anonym{Anonym: pbAuth},
		}
		return request, nil
	case *auth.QBasicAuth:
		pbAuth, err := v.Pb()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		request := &proto.QPBxAuthRequest{
			Auth: &proto.QPBxAuthRequest_Basic{Basic: pbAuth},
		}
		return request, nil
	default:
		err := errors.New("Unknown object type")
		msg := fmt.Sprintf("%T", v)
		log.Error().Stack().Err(err).Msg(msg)
		return nil, errors.WithStack(err)
	}
}
