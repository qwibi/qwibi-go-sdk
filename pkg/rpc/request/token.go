package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QTokenRequest struct {
	RequestId string
	LayerId   string
}

func NewTokenRequest(layerId string) (*QTokenRequest, error) {
	if layerId == "" {
		return nil, qlog.Error("layer ID is not defined")
	}

	req := &QTokenRequest{
		RequestId: utils.RequestId(),
		LayerId:   layerId,
	}

	return req, nil
}

func NewTokenRequestPb(in *proto.QPBxTokenRequest) (*QTokenRequest, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	return NewTokenRequest(in.LayerId)
}

func (c *QTokenRequest) Pb() *proto.QPBxTokenRequest {
	return &proto.QPBxTokenRequest{
		RequestId: c.RequestId,
		LayerId:   c.LayerId,
	}
}

func (c *QTokenRequest) Message() protobuf.Message {
	return c.Pb()
}

func (c *QTokenRequest) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QTokenRequest) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_TokenRequest{
			TokenRequest: c.Pb(),
		},
	}
}
