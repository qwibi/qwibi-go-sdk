package geojson

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QPoint ...
type QPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (c *QPoint) Valid() error {
	if c.Type == "" || len(c.Coordinates) < 2 {
		errors.New("Invalid point format")
	}

	return nil
}

// GetType ...
func (c *QPoint) GetType() string {
	return c.Type
}

// Pb ...
func (c *QPoint) Pb() *proto.QPBxPoint {
	pb := &proto.QPBxPoint{Coordinates: c.Coordinates}
	return pb
}

// GeometryPb ...
func (c *QPoint) GeometryPb() (*proto.QPBxGeometry, error) {
	if err := c.Valid(); err != nil {
		return nil, errors.WithStack(err)
	}

	pb := &proto.QPBxGeometry{Object: &proto.QPBxGeometry_Point{Point: c.Pb()}}

	return pb, nil
}

// SetCenter ...
func (c *QPoint) SetCenter(x float64, y float64) {
	c.SetCoordinates(x, y)
}

// SetCoordinates ...
func (c *QPoint) SetCoordinates(x float64, y float64) {
	c.Coordinates = []float64{x, y}
}

// NewPoint ...
func NewPoint(coordinates []float64) (*QPoint, error) {
	p := &QPoint{
		Type:        QPointType,
		Coordinates: coordinates}

	return p, p.Valid()
}

// NewPointLatLng ...
func NewPointLatLng(lat float64, lng float64) (*QPoint, error) {
	return NewPoint([]float64{lng, lat})
}

// NewPointZero ...
func NewPointZero() *QPoint {
	point, _ := NewPoint([]float64{0, 0})
	return point
}

// NewPointPb ...
func NewPointPb(pb *proto.QPBxPoint) (*QPoint, error) {
	if pb == nil {
		return nil, errors.New("Invalid format type nil")
	}

	return NewPoint(pb.Coordinates)
}

// NewPointGeometryPb ...
func NewPointGeometryPb(pb *proto.QPBxGeometry_Point) (*QPoint, error) {
	if pb == nil {
		return nil, errors.New("Invalid format type nil")
	}

	return NewPointPb(pb.Point)
}
