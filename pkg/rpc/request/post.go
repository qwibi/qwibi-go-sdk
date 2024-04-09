package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QPostRequest struct {
	RequestId string
	LayerId   string
	Object    *geo.QGeoObject
}

func NewPostRequest(layerId string, object *geo.QGeoObject) (*QPostRequest, error) {
	if layerId == "" {
		return nil, qlog.Error("layer ID is not defined")
	}

	if object == nil {
		return nil, qlog.Error("object is not defined")
	}

	req := &QPostRequest{
		RequestId: utils.RequestId(),
		LayerId:   layerId,
		Object:    object,
	}

	return req, nil
}

func NewPostRequestPb(in *proto.QPBxPostRequest) (*QPostRequest, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	object, err := geo.NewGeoObjectPb(in.Object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewPostRequest(in.LayerId, object)
}

func (c *QPostRequest) Pb() *proto.QPBxPostRequest {
	return &proto.QPBxPostRequest{
		RequestId: c.RequestId,
		LayerId:   c.LayerId,
		Object:    c.Object.Pb(),
	}
}

func (c *QPostRequest) Message() protobuf.Message {
	return c.Pb()
}

func (c *QPostRequest) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QPostRequest) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_PostRequest{
			PostRequest: c.Pb(),
		},
	}
}
