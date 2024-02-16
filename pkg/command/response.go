package command

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QResponse struct {
	Status string
}

func NewResponse(status string) *QResponse {
	return &QResponse{
		Status: status,
	}
}

func NewResponsePb(in *proto.QPBxResponse) (*QResponse, error) {
	if in == nil {
		return nil, qlog.Error("wrong parameter type nil")
	}

	res := &QResponse{
		Status: in.Status,
	}

	return res, nil
}

func (c *QResponse) Pb() *proto.QPBxResponse {
	return &proto.QPBxResponse{
		Status: c.Status,
	}
}
