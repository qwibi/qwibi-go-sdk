package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/query"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QQueryRequest struct {
	RequestId string
	LayerId   string
}

func NewQueryRequest(layerId string, options ...query.QueryOption) (*QQueryRequest, error) {
	req := &QQueryRequest{
		RequestId: utils.RequestId(),
		LayerId:   layerId,
	}

	return req, nil
}

func NewQueryRequestPb(in *proto.QPBxQueryRequest) (*QQueryRequest, error) {
	return NewQueryRequest(
		in.LayerId,
	)
}

func (c *QQueryRequest) Pb() *proto.QPBxQueryRequest {
	return &proto.QPBxQueryRequest{
		RequestId: c.RequestId,
		LayerId:   c.LayerId,
	}
}
