package event

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QEvent interface {
	Valid() error
	Pb() *proto.QPBxEvent
}

// NewGeometryPb ...
func NewEventPb(in *proto.QPBxEvent) (QEvent, error) {
	if in == nil {
		return nil, qlog.Error("Event: bed parameter type nil")
	}

	switch v := in.Type.(type) {
	case *proto.QPBxEvent_LayerUpdate:
		return NewLayerUpdatePb(v.LayerUpdate)
	case *proto.QPBxEvent_ObjectUpdate:
		return NewObjectUpdatePb(v.ObjectUpdate)
	default:
		return nil, qlog.Errorf("Event: unknown event type: %s", v)
	}
}
