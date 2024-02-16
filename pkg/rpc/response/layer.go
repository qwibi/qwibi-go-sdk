package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QLayerResponse struct {
	reqId string
	Layer *layer.QGeoLayer
}

func NewLayerResponse(layer *layer.QGeoLayer) (*QLayerResponse, error) {
	r := &QLayerResponse{
		Layer: layer,
	}

	return r, nil
}

func NewLayerResponsePb(in *proto.QPBxLayerResponse) (*QLayerResponse, error) {
	layer, err := layer.NewGeoLayerPb(in.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewLayerResponse(layer)
}

func (c *QLayerResponse) Pb() *proto.QPBxLayerResponse {
	return &proto.QPBxLayerResponse{
		Layer: c.Layer.Pb(),
	}
}

func (c *QLayerResponse) ReqId() string {
	return c.reqId
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
