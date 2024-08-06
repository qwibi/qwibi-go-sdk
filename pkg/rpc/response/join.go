package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QJoinResponse struct {
	RequestId string
	Layer     *layer.QLayer
}

func NewJoinResponse(requestId string, layer *layer.QLayer) (*QJoinResponse, error) {
	res := &QJoinResponse{
		RequestId: requestId,
		Layer:     layer,
	}

	return res, res.Valid()
}

func NewJoinResponsePb(in *proto.QPBxJoinResponse) (*QJoinResponse, error) {
	layer, err := layer.NewGeoLayerPb(in.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewJoinResponse(in.RequestId, layer)
}

func (c *QJoinResponse) Pb() *proto.QPBxJoinResponse {
	return &proto.QPBxJoinResponse{
		RequestId: c.RequestId,
		Layer:     c.Layer.Pb(),
	}
}

func (c *QJoinResponse) Valid() error {
	if c.RequestId == "" {
		return qlog.Error("request ID is not defined")
	}

	if c.Layer == nil {
		return qlog.Error("layer is not defined")
	}

	return nil
}
