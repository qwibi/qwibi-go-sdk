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

	req := &QPostRequest{
		RequestId: utils.RequestId(),
		LayerId:   layerId,
		Object:    object,
	}

	return req, req.Valid()
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
		Object:    c.Object.Pb(),
	}
}

func (c *QPostRequest) Valid() error {
	if c.RequestId == "" {
		return qlog.Error("request ID is not defined")
	}

	if c.LayerId == "" {
		return qlog.Error("layer ID is not defined")
	}

	if c.Object == nil {
		return qlog.Error("object is not defined" +
			"")
	}

	return nil
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
