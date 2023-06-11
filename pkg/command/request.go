package command

import "github.com/qwibi/qwibi-go-sdk/proto"

type QRequest struct {
	Path string
	Args []map[string]string
}

func NewCRequest(path string) QRequest {
	return QRequest{Path: path}
}

func (c *QRequest) Pb() *proto.QPBxCommandRequest {
	return &proto.QPBxCommandRequest{
		Command: c.Path,
	}
}
