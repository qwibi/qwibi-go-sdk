package command

import "github.com/qwibi/qwibi-go-sdk/proto"

type QResponse struct {
	Path string
	Args []map[string]string
}

func NewResponse(path string) QResponse {
	return QResponse{Path: path}
}

func (c *QResponse) Pb() *proto.QPBxCommandRequest {
	return &proto.QPBxCommandRequest{
		Command: c.Path,
	}
}
