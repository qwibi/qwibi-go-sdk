package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QAuthRequest struct {
	RequestId string
	Auth      auth.QAuth
}

func NewAuthRequest(auth auth.QAuth) (*QAuthRequest, error) {
	req := &QAuthRequest{
		RequestId: utils.RequestId(),
		Auth:      auth,
	}

	return req, req.Valid()
}

func NewAuthRequestPb(in *proto.QPBxAuthRequest) (*QAuthRequest, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	auth, err := auth.NewAuthPb(in.Auth)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewAuthRequest(auth)
}

func (c *QAuthRequest) Pb() *proto.QPBxAuthRequest {
	return &proto.QPBxAuthRequest{
		RequestId: c.RequestId,
		Auth:      c.Auth.Pb(),
	}
}

func (c *QAuthRequest) Valid() error {
	if c.RequestId == "" {
		return qlog.Error("request ID is not defined")
	}

	if c.Auth == nil {
		return qlog.Error("auth method is not defined")
	}

	return nil
}
