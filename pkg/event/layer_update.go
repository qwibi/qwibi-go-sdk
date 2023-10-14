package event

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPoint ...
type QLayerUpdate struct {
	Layer *layer.QGeoLayer `json:"object"`
}

// UpdateEvent ...
func NewLayerUpdate(object *layer.QGeoLayer) (*QLayerUpdate, error) {
	if object == nil {
		return nil, qlog.Error("Bed parameter type nil")
	}

	event := &QLayerUpdate{
		Layer: object,
	}

	return event, nil
}

// NewUpdateEventPb ...
func NewLayerUpdatePb(pb *proto.QPBxLayerUpdate) (*QLayerUpdate, error) {
	if pb == nil {
		return nil, qlog.Error("LayerUpdate: bed parameter type nil")
	}

	object, err := layer.NewGeoLayerPb(pb.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewLayerUpdate(object)
}

// Valid ...
func (c *QLayerUpdate) Valid() error {
	if c.Layer == nil {
		return qlog.Error("LayerUpdate: object not defined")
	}

	return nil
}

// Pb ...
func (c *QLayerUpdate) Pb() *proto.QPBxEvent {
	if c.Layer == nil {
		qlog.Error("LayerUpdate: object is not defined")
		return nil
	}

	pb := &proto.QPBxEvent{
		Type: &proto.QPBxEvent_LayerUpdate{
			LayerUpdate: &proto.QPBxLayerUpdate{
				Layer: c.Layer.Pb(),
			},
		},
	}

	return pb
}
