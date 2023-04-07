package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/feature"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
	"github.com/qwibi/qwibi-go-sdk/utils"
)

type QGeoPoint struct {
	gid        string
	Feature    *feature.QPointFeature
	Properties []byte
}

// NewGeoPoint ...
func NewGeoPoint() *QGeoPoint {
	return &QGeoPoint{
		gid:     utils.NewID(),
		Feature: feature.NewPointFeature(),
	}
}

func NewGeoPointPb(pb *proto.QPBxGeoObject) (*QGeoPoint, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	pointPb := pb.GetPoint()
	if pointPb == nil {
		err := errors.New("Invalid GeoLayer format type nil")
		return nil, qlog.Error(err)
	}

	featurePb, err := feature.NewPointFeaturePb(pointPb.Feature)
	if err != nil {
		return nil, qlog.Error(err)
	}

	point := &QGeoPoint{
		gid:        pointPb.Gid,
		Feature:    featurePb,
		Properties: pointPb.Properties,
	}

	return point, point.Valid()
}

func (c *QGeoPoint) Valid() error {
	return c.Feature.Geometry.Valid()
}

func (c *QGeoPoint) Gid() string {
	return c.gid
}

func (c *QGeoPoint) Pb() *proto.QPBxGeoObject {
	return &proto.QPBxGeoObject{
		Geo: &proto.QPBxGeoObject_Point{
			Point: &proto.QPBxGeoPoint{
				Gid:        c.gid,
				Feature:    c.Feature.Pb(),
				Properties: c.Properties,
			},
		},
	}
}
