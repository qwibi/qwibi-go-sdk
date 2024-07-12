package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QTokenResponse struct {
	RequestId string
	Token     string
}

func NewTokenResponse(requestId string, token string) (*QTokenResponse, error) {
	if requestId == "" {
		return nil, qlog.Error("request ID not defined")
	}

	if token == "" {
		return nil, qlog.Error("token ID is not defined")
	}

	res := &QTokenResponse{
		RequestId: requestId,
		Token:     token,
	}

	return res, nil
}

func NewTokenResponsePb(in *proto.QPBxTokenResponse) (*QTokenResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	return NewTokenResponse(in.RequestId, in.Token)
}

func (c *QTokenResponse) Pb() *proto.QPBxTokenResponse {
	return &proto.QPBxTokenResponse{
		RequestId: c.RequestId,
		Token:     c.Token,
	}
}

func (c *QTokenResponse) Message() protobuf.Message {
	return c.Pb()
}

func (c *QTokenResponse) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QTokenResponse) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_TokenResponse{
			TokenResponse: c.Pb(),
		},
	}
}
