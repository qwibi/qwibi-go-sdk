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

func NewGeoObject(g Object) *QGeoObject {
	return &QGeoObject{g}
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
	switch v := in.Geometry.Type.(type) {
	case *proto.QPBxGeometry_Point:
		geometry, err := geometry.NewPointPb(v.Point)
		if err != nil {
			return nil, qlog.Error(err)
		}
		point := NewGeoPoint(
			PointGid(in.Gid),
			PointCoordinates(geometry.Coordinates...),
			PointProperties(in.Properties),
		)
		return NewGeoObject(point), nil
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
