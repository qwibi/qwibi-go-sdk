package event

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/object"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPoint ...
type QObjectUpdate struct {
	Object *object.QGeoObject `json:"object"`
}

// UpdateEvent ...
func NewObjectUpdate(object *object.QGeoObject) *QObjectUpdate {
	return &QObjectUpdate{
		Object: object,
	}
}

// NewUpdateEventPb ...
func NewObjectUpdatePb(pb *proto.QPBxObjectUpdate) (*QObjectUpdate, error) {
	if pb == nil {
		return nil, qlog.Error("ObjectUpdate: bed parameter type nil")
	}

	object, err := object.NewGeoObjectPb(pb.Object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	event := &QObjectUpdate{
		Object: object,
	}

	return event, event.Valid()
}

// Valid ...
func (c *QObjectUpdate) Valid() error {
	if c.Object == nil {
		return qlog.Error("ObjectUpdate: object not defined")
	}

	return c.Object.Valid()
}

// Pb ...
func (c *QObjectUpdate) Pb() *proto.QPBxEvent {
	if c.Object != nil {
		qlog.Error("ObjectUpdate: object is not defined")
		return nil
	}

	pb := &proto.QPBxEvent{
		Type: &proto.QPBxEvent_Update{
			Update: &proto.QPBxObjectUpdate{
				Object: c.Object.Pb(),
			},
		},
	}

	return pb
}
