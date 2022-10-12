package geometry

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
)

// QPoint ...
type QPoint struct {
	Coordinates []float64 `json:"coordinates"`
}

// NewPoint ...
func NewPoint() *QPoint {
	return &QPoint{
		Coordinates: []float64{0, 0},
	}
}

// NewPointPb ...
func NewPointPb(pb *proto.QPBxPoint) (*QPoint, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	point := NewPoint()
	point.Coordinates = pb.Coordinates

	return point, point.Valid()
}

// Valid ...
func (c *QPoint) Valid() error {
	if len(c.Coordinates) < 2 {
		err := errors.New("Wrong coordinates format")
		return qlog.Error(err)
	}

	return nil
}

// Pb ...
func (c *QPoint) Pb() *proto.QPBxPoint {
	return &proto.QPBxPoint{
		Coordinates: c.Coordinates,
	}
}
