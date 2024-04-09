package command

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"strings"
)

type QCommand struct {
	Command string
	Args    []string
}

func NewCommand(command string, args ...string) (*QCommand, error) {
	if command == "" {
		return nil, qlog.Error("command not defined")
	}

	req := &QCommand{
		Command: command,
		Args:    args,
	}

	return req, req.Valid()
}

func NewCommandPb(in *proto.QPBxCommand) (*QCommand, error) {
	if in == nil {
		return nil, qlog.Error("command request not defined")
	}

	return NewCommand(in.Command, in.Args...)
}

func (c *QCommand) Pb() *proto.QPBxCommand {
	return &proto.QPBxCommand{
		Command: c.Command,
		Args:    c.Args,
	}
}

func (c *QCommand) Valid() error {
	if c.Command == "" {
		return qlog.Error("command not defined")
	}

	if !strings.HasPrefix(c.Command, "/") {
		return qlog.Error("command must have '/' prefix")
	}

	return nil
}
