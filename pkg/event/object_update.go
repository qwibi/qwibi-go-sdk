package event

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPoint ...
type QObjectUpdate struct {
	Object *geo.QGeoObject `json:"object"`
}

// UpdateEvent ...
func NewObjectUpdate(object *geo.QGeoObject) (*QObjectUpdate, error) {
	if object == nil {
		return nil, qlog.Error("Bed parameter type nil")
	}

	event := &QObjectUpdate{
		Object: object,
	}

	return event, nil
}

// NewUpdateEventPb ...
func NewObjectUpdatePb(pb *proto.QPBxObjectUpdate) (*QObjectUpdate, error) {
	if pb == nil {
		return nil, qlog.Error("ObjectUpdate: bed parameter type nil")
	}

	object, err := geo.NewGeoObjectPb(pb.Object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewObjectUpdate(object)
}

// Valid ...
func (c *QObjectUpdate) Valid() error {
	if c.Object == nil {
		return qlog.Error("ObjectUpdate: object not defined")
	}

	return nil
}

// Pb ...
func (c *QObjectUpdate) Pb() *proto.QPBxEvent {
	if c.Object == nil {
		qlog.Error("ObjectUpdate: object is not defined")
		return nil
	}

	pb := &proto.QPBxEvent{
		Type: &proto.QPBxEvent_ObjectUpdate{
			ObjectUpdate: &proto.QPBxObjectUpdate{
				Object: c.Object.Pb(),
			},
		},
	}

	return pb
}
