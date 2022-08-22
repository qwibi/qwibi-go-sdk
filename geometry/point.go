package geometry

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
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
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	point := NewPoint()
	point.Coordinates = pb.Coordinates

	return point, point.Valid()
}

// Valid ...
func (c *QPoint) Valid() error {
	if len(c.Coordinates) < 2 {
		err := errors.New("Wrong coordinates format")
		log.Error().Stack().Err(err).Msg("")
		return err
	}

	return nil
}

// Pb ...
func (c *QPoint) Pb() *proto.QPBxPoint {
	return &proto.QPBxPoint{
		Coordinates: c.Coordinates,
	}
}
