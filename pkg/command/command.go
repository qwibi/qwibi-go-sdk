package command

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QCommand struct {
	name        string `json:"name"`
	description string `json:"description"`
}

func NewCommand(name string, description string) (*QCommand, error) {
	if name == "" {
		return nil, qlog.Error("command name not defined")
	}

	command := &QCommand{
		name:        name,
		description: description,
	}

	return command, nil
}

func NewCommandPb(in *proto.QPBxCommand) (*QCommand, error) {
	if in == nil {
		return nil, qlog.Error("command not defined")
	}

	command := &QCommand{
		name:        in.Name,
		description: in.Description,
	}

	return command, nil
}

// Pb ...
func (c *QCommand) Pb() *proto.QPBxCommand {
	pb := &proto.QPBxCommand{
		Name:        c.name,
		Description: c.description,
	}

	return pb
}
