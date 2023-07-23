package geometry

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeometry ...
type QGeometry interface {
	Valid() error
	Pb() *proto.QPBxGeometry
	GetType() string
	//UnmarshalJSON(data []byte) error
}

// NewGeometryPb ...
func NewGeometryPb(in *proto.QPBxGeometry) (QGeometry, error) {
	if in == nil {
		return nil, qlog.Error("Wrong geometry parameter type nil")
	}

	switch v := in.Type.(type) {
	case *proto.QPBxGeometry_Point:
		return NewPointPb(v.Point)
	default:
		return nil, qlog.Errorf("Unknown geometry type: %s", v)
	}
}
