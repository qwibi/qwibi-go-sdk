package request

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/auth"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QAuthRequest ...
type QAuthRequest struct {
	Auth interface{}
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
		return nil, errors.New("Unknown auth request")
	}
}

// NewAuthRequest ...
func NewAuthRequest(authMethod interface{}) (*QAuthRequest, error) {
	switch t := authMethod.(type) {
	case *auth.QAnonymousAuth:
		req := &QAuthRequest{
			Auth: t,
		}
		return req, nil
	case *auth.QBasicAuth:
		req := &QAuthRequest{}
		return req, nil
	default:
		return nil, errors.New("Unknown auth request")
	}
}

// NewAuthRequestPb ...
func NewAuthRequestPb(pb *proto.QPBxAuthRequest) (*QAuthRequest, error) {
	switch v := pb.Auth.(type) {
	case *proto.QPBxAuthRequest_Anonym:
		t, err := auth.NewAnonymousAuthPb(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &QAuthRequest{Auth: t}, nil
	case *proto.QPBxAuthRequest_Basic:
		t, err := auth.NewBasicAuthPb(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &QAuthRequest{Auth: t}, nil
	default:
		return nil, errors.New("Unknown auth request")
	}
}
