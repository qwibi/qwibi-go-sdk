package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QPostResponse struct {
	RequestId string
	Gid       string
}

func NewPostResponse(requestId string, gid string) (*QPostResponse, error) {
	if requestId == "" {
		return nil, qlog.Error("request ID not defined")
	}

	if gid == "" {
		return nil, qlog.Error("gid ID is not defined")
	}

	res := &QPostResponse{
		RequestId: requestId,
		Gid:       gid,
	}

	return res, nil
}

func NewPostResponsePb(in *proto.QPBxPostResponse) (*QPostResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	return NewPostResponse(in.RequestId, in.Gid)
}

func (c *QPostResponse) Pb() *proto.QPBxPostResponse {
	return &proto.QPBxPostResponse{
		Gid: c.Gid,
	}
}

func (c *QPostResponse) Message() protobuf.Message {
	return c.Pb()
}

func (c *QPostResponse) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QPostResponse) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_PostResponse{
			PostResponse: c.Pb(),
		},
	}
}
