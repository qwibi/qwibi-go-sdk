package rpc

import (
	"fmt"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/rpc/request"
	"github.com/qwibi/qwibi-go-sdk/pkg/rpc/response"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QEvent interface {
	//reqId() string
	ProtoReflect() protoreflect.Message
	Event() *proto.QPBxEvent
	//Message() protobuf.Message
}

//func (c QEvent) ProtoReflect() protoreflect.Message {
//	//TODO implement me
//	panic("implement me")
//}

//func NewMessage(m QEvent) (*QEvent, error) {
//	switch v := m.(type) {
//	case *request.QLayerRequest:
//		r := &QEvent{
//			: v,
//		}
//
//		return r, nil
//	default:
//		return nil, qlog.Error("unknown event type")
//	}
//}

func NewEventPb(in *proto.QPBxEvent) (QEvent, error) {
	if in == nil {
		return nil, qlog.Error("bed parameter type nil")
	}

	switch v := in.Type.(type) {
	case *proto.QPBxEvent_LayerRequest:
		return request.NewLayerRequestPb(v.LayerRequest)
	case *proto.QPBxEvent_LayerResponse:
		return response.NewLayerResponsePb(v.LayerResponse)
	case *proto.QPBxEvent_CommandRequest:
		return request.NewCommandRequestPb(v.CommandRequest)
	case *proto.QPBxEvent_CommandResponse:
		return response.NewCommandResponsePb(v.CommandResponse)
	default:
		return nil, fmt.Errorf("unknown message type: %T", v)
	}
}

func NewMessageBytes(b []byte) (QEvent, error) {

	return nil, nil
}
