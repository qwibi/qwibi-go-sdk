package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QPostResponse struct {
	RequestId string
	LayerId   string
	Object    *geo.QGeoObject
}

func NewPostResponse(requestId string, layerId string, object *geo.QGeoObject) (*QPostResponse, error) {
	res := &QPostResponse{
		RequestId: requestId,
		LayerId:   layerId,
		Object:    object,
	}

	return res, nil
}

func NewPostResponsePb(in *proto.QPBxPostResponse) (*QPostResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	object, err := geo.NewGeoObjectPb(in.Object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewPostResponse(in.RequestId, in.LayerId, object)
}

func (c *QPostResponse) Pb() *proto.QPBxPostResponse {
	pb := &proto.QPBxPostResponse{
		RequestId: c.RequestId,
		LayerId:   c.LayerId,
		Object:    c.Object.Pb(),
	}

	return pb
}

func (c *QPostResponse) Valid() error {
	if c.RequestId == "" {
		return qlog.Error("request ID is not defined")
	}

	if c.Object == nil {
		return qlog.Error("object is not defined")
	}

	if c.Object.Valid() != nil {
		return qlog.Error("invalid object")
	}

	return nil
}

func (c *QPostResponse) Message() protobuf.Message {
	return c.Pb()
}

func (c *QPostResponse) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QPostResponse) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_PostResponse{
			PostResponse: c.Pb(),
		},
	}
}
