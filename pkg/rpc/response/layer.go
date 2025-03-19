package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QLayerResponse struct {
	RequestId string
	Layer     *layer.QLayer
}

func NewLayerResponse(requestId string, layer *layer.QLayer) (*QLayerResponse, error) {
	if requestId == "" {
		return nil, qlog.Error("request ID not defined")
	}

	if layer == nil {
		return nil, qlog.Error("layer is not defined")
	}

	res := &QLayerResponse{
		RequestId: requestId,
		Layer:     layer,
	}

	return res, nil
}

func NewLayerResponsePb(in *proto.QPBxLayerResponse) (*QLayerResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	layer, err := layer.NewLayerPb(in.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewLayerResponse(in.RequestId, layer)
}

func (c *QLayerResponse) Pb() *proto.QPBxLayerResponse {
	return &proto.QPBxLayerResponse{
		RequestId: c.RequestId,
		Layer:     c.Layer.Pb(),
	}
}

func (c *QLayerResponse) Message() protobuf.Message {
	return c.Pb()
}

func (c *QLayerResponse) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QLayerResponse) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_LayerResponse{
			LayerResponse: c.Pb(),
		},
	}
}
