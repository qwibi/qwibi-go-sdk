package geometry

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

// QPoint ...
type QPoint struct {
	Coordinates []float64 `json:"coordinates"`
	Properties  *structpb.Struct
}

// NewPoint ...
func NewPoint(coordinates ...float64) (*QPoint, error) {
	p := &QPoint{
		Coordinates: coordinates,
		Properties:  &structpb.Struct{},
	}

	return p, p.Valid()
}

// NewZeroPoint ...
func NewZeroPoint() *QPoint {
	point, _ := NewPoint(0, 0)
	return point
}

// NewPointPb ...
func NewPointPb(pb *proto.QPBxPoint) (*QPoint, error) {
	if pb == nil {
		err := errors.New("Invalid point format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	point, err := NewPoint(pb.Coordinates...)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	if pb.Properties != nil {
		point.Properties = pb.Properties
	}

	return point, nil
}

// SetCenter ...
func (c *QPoint) SetCenter(x float64, y float64) {
	c.SetCoordinates(x, y)
}

// SetCoordinates ...
func (c *QPoint) SetCoordinates(x float64, y float64) {
	c.Coordinates = []float64{x, y}
}

// Valid ...
func (c *QPoint) Valid() error {
	if len(c.Coordinates) < 2 {
		errors.New("Invalid point format")
	}

	return nil
}

// Pb ...
func (c *QPoint) Pb() (*proto.QPBxGeometry, error) {
	if len(c.Coordinates) < 2 {
		err := errors.New("Wrong coordinates format")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxGeometry{
		Geometry: &proto.QPBxGeometry_Point{
			Point: &proto.QPBxPoint{Coordinates: c.Coordinates},
		},
	}

	return pb, nil
}
