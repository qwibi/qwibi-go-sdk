package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QJoinRequest struct {
	RequestId string
	LayerId   string
}

func NewJoinRequest(layerId string) (*QJoinRequest, error) {
	req := &QJoinRequest{
		RequestId: utils.RequestId(),
		LayerId:   layerId,
	}

	return req, req.Valid()
}

func NewJoinRequestPb(in *proto.QPBxJoinRequest) (*QJoinRequest, error) {
	return NewJoinRequest(in.LayerId)
}

func (c *QJoinRequest) Pb() *proto.QPBxJoinRequest {
	return &proto.QPBxJoinRequest{
		LayerId: c.LayerId,
	}
}

func (c *QJoinRequest) Valid() error {
	if c.LayerId == "" {
		return qlog.Error("layer ID is not defined")
	}

	return nil
}
