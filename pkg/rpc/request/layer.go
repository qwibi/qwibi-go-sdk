package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QLayerRequest struct {
	RequestId string
	Layer     *layer.QLayer
}

func NewLayerRequest(options ...layer.LayerOption) (*QLayerRequest, error) {
	layer := layer.NewLayer(options...)

	req := &QLayerRequest{
		RequestId: utils.RequestId(),
		Layer:     layer,
	}

	return req, layer.Valid()
}

func NewLayerRequestPb(in *proto.QPBxLayerRequest) (*QLayerRequest, error) {
	return NewLayerRequest(
		layer.WithLayerGid(in.LayerId),
		layer.WithLayerName(in.Name),
		layer.WithLayerPublic(in.Public),
		layer.WithLayerProperties(in.Properties),
		layer.WithLayerCommands(in.Commands),
	)
}

func (c *QLayerRequest) Pb() *proto.QPBxLayerRequest {
	return &proto.QPBxLayerRequest{
		LayerId:    c.Layer.LayerId,
		Name:       c.Layer.Name,
		Public:     c.Layer.Public,
		Properties: c.Layer.Properties,
		Commands:   c.Layer.Commands,
	}
}

func (c *QLayerRequest) Message() protobuf.Message {
	return c.Pb()
}

func (c *QLayerRequest) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QLayerRequest) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_LayerRequest{
			LayerRequest: c.Pb(),
		},
	}
}
