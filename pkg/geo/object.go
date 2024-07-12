package geo

import (
	"encoding/json"
	sdkGeometry "github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeometry ...
type Object interface {
	Gid() string
	Geometry() *sdkGeometry.QGeometry
	Properties() []byte
	Pb() *proto.QPBxGeoObject
	Valid() error
}

type QGeoObject struct {
	Object
}

func GeoObject(g Object) *QGeoObject {
	return &QGeoObject{g}
}

func NewGeoObject(gid string, geometry *sdkGeometry.QGeometry, properties []byte) (*QGeoObject, error) {
	if geometry == nil {
		return nil, qlog.Error("geometry not defined")
	}

	switch v := geometry.Geometry.(type) {
	case *sdkGeometry.QPoint:
		point := NewGeoPoint(
			WithPointGid(gid),
			WithPointGeometry(v),
			WithPointProperties(properties),
		)
		return GeoObject(point), point.Valid()
	default:
		return nil, qlog.Error("unknown geometry type %T", v)
	}
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
	switch v := in.Geometry.Type.(type) {
	case *proto.QPBxGeometry_Point:
		geometry, err := sdkGeometry.NewPointPb(v.Point)
		if err != nil {
			return nil, qlog.Error(err)
		}

		point := NewGeoPoint(
			WithPointCoordinates(geometry.Coordinates...),
			WithPointProperties(in.Properties),
		)

		return GeoObject(point), point.Valid()
	default:
		return nil, qlog.Errorf("Unknown geometry type: %T", v)
	}

}

func NewGeoObjectBytes(data []byte) (*QGeoObject, error) {
	var object QGeoObject

	err := json.Unmarshal(data, &object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return &object, nil
}
