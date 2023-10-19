package geo

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeometry ...
type Object interface {
	Pb() *proto.QPBxGeoObject
	Gid() string
	Geometry() *geometry.QGeometry
	Properties() []byte
	Valid() error
}

type QGeoObject struct {
	Object
}

func GeoObject(g Object) *QGeoObject {
	return &QGeoObject{g}
}

func NewGeoObject(gid string, geom *geometry.QGeometry, properties []byte) (*QGeoObject, error) {
	if geom == nil {
		return nil, qlog.Error("geometry not defined")
	}

	switch v := geom.Geometry.(type) {
	case *geometry.QPoint:
		g := NewGeoPoint(
			WithPointGid(gid),
			WithPointGeometry(v),
			WithPointProperties(properties),
		)
		return &QGeoObject{g}, nil
	default:
		return nil, qlog.Error("unknown geometry type %T", v)
	}
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
	switch v := in.Geometry.Type.(type) {
	case *proto.QPBxGeometry_Point:
		geometry, err := geometry.NewPointPb(v.Point)
		if err != nil {
			return nil, qlog.Error(err)
		}
		point := NewGeoPoint(
			WithPointGid(in.Gid),
			WithPointCoordinates(geometry.Coordinates...),
			WithPointProperties(in.Properties),
		)
		return GeoObject(point), nil
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
