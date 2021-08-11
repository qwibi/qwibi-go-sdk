package geojson

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeometry ...
type QGeometry interface {
	GetType() string
	GeometryPb() (*proto.QPBxGeometry, error)
	SetCenter(x float64, y float64)
	Valid() error
}

// NewGeometryPb ...
func NewGeometryPb(pb *proto.QPBxGeometry) (QGeometry, error) {
	if pb == nil || pb.Object == nil {
		return nil, errors.New("Invalid format type nil")
	}

	switch v := pb.Object.(type) {
	case *proto.QPBxGeometry_Point:
		return NewPointPb(v.Point)
	default:
		return nil, errors.New("Unknown geojson type")
	}
}
