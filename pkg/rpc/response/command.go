package response

import (
	sdkCommand "github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QCommandResponse struct {
	RequestId string
	LayerId   string
	Response  *sdkCommand.QResponse
}

func NewCommandResponse(requestId string, layerId string, response *sdkCommand.QResponse) (*QCommandResponse, error) {
	if requestId == "" {
		return nil, qlog.Error("layer ID is not defined")
	}

	if layerId == "" {
		return nil, qlog.Error("layer ID is not define")
	}

	if response == nil {
		return nil, qlog.Error("response is not defined")
	}

	r := &QCommandResponse{
		RequestId: requestId,
		LayerId:   layerId,
		Response:  response,
	}

	return r, nil
}

func NewCommandResponsePb(in *proto.QPBxCommandResponse) (*QCommandResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	response, err := sdkCommand.NewResponsePb(in.Response)
	if err != nil {
		return nil, qlog.Error(err)
	}

	res, err := NewCommandResponse(in.RequestId, in.LayerId, response)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return res, nil
}

func (c *QCommandResponse) Pb() *proto.QPBxCommandResponse {
	return &proto.QPBxCommandResponse{
		RequestId: c.RequestId,
		LayerId:   c.LayerId,
		Response:  c.Response.Pb(),
	}
}

//func (c *QCommandResponse) ReqId() string {
//	return c.RequestId
//}

func (c *QCommandResponse) Message() protobuf.Message {
	return c.Pb()
}

func (c *QCommandResponse) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QCommandResponse) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_CommandResponse{
			CommandResponse: c.Pb(),
		},
	}
}
