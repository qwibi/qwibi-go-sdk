package request

import (
	sdkCommand "github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

type QCommandRequest struct {
	RequestId string
	LayerId   string
	Command   *sdkCommand.QCommand
}

func NewCommandRequest(layerId string, command *sdkCommand.QCommand) (*QCommandRequest, error) {
	req := &QCommandRequest{
		RequestId: utils.RequestId(),
		LayerId:   strings.TrimSpace(layerId),
		Command:   command,
	}

	return req, nil
}

func NewCommandRequestPb(in *proto.QPBxCommandRequest) (*QCommandRequest, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	command, err := sdkCommand.NewCommandPb(in.Command)
	if err != nil {
		return nil, qlog.Error(err)
	}

	req := &QCommandRequest{
		RequestId: in.RequestId,
		LayerId:   in.LayerId,
		Command:   command,
	}

	return req, nil
}

func (c *QCommandRequest) Pb() *proto.QPBxCommandRequest {
	return &proto.QPBxCommandRequest{
		RequestId: c.RequestId,
		LayerId:   c.LayerId,
		Command:   c.Command.Pb(),
	}
}

func (c *QCommandRequest) Message() protobuf.Message {
	return c.Pb()
}

func (c *QCommandRequest) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QCommandRequest) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_CommandRequest{
			CommandRequest: c.Pb(),
		},
	}
}
