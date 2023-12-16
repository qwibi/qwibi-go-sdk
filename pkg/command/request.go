package command

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QRequest struct {
	command *QCommand
}

func NewRequest(command *QCommand) (*QRequest, error) {
	if command == nil {
		return nil, qlog.Error("command not defined")
	}
	return &QRequest{command: command}, nil
}

func NewRequestPb(in *proto.QPBxCommandRequest) (*QRequest, error) {
	if in == nil {
		return nil, qlog.Error("command request not defined")
	}

	command, err := NewCommandPb(in.Command)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewRequest(command)
}

func (c *QRequest) Pb() *proto.QPBxCommandRequest {
	return &proto.QPBxCommandRequest{
		Command: c.command.Pb(),
	}
}
