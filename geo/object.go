package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
)

type QGeoObject interface {
	Valid() error
	Pb() *proto.QPBxGeoObject
}

// NewGeoObjectPb ...
func NewGeoObjectPb(pb *proto.QPBxGeoObject) (QGeoObject, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	switch v := pb.Geo.(type) {
	case *proto.QPBxGeoObject_Layer:
		return NewGeoLayerPb(pb)
	case *proto.QPBxGeoObject_Point:
		return NewGeoPointPb(pb)
	default:
		return nil, qlog.Error("Unknown geometry type: %T", v)
	}
}
