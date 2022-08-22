package geo

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/feature"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
)

type QGeoPoint struct {
	Gid        string
	Feature    *feature.QPointFeature
	Properties []byte
}

// NewGeoPoint ...
func NewGeoPoint() *QGeoPoint {
	return &QGeoPoint{
		Gid:     utils.NewID(),
		Feature: feature.NewPointFeature(),
	}
}

func NewGeoPointPb(pb *proto.QPBxGeoObject) (*QGeoPoint, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	pointPb := pb.GetPoint()
	if pointPb == nil {
		err := errors.New("Invalid GeoLayer format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	featurePb, err := feature.NewPointFeaturePb(pointPb.Feature)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	point := &QGeoPoint{
		Gid:        pointPb.Gid,
		Feature:    featurePb,
		Properties: pointPb.Properties,
	}

	return point, point.Valid()
}

func (c *QGeoPoint) Valid() error {
	return c.Feature.Geometry.Valid()
}

func (c *QGeoPoint) Pb() *proto.QPBxGeoObject {
	return &proto.QPBxGeoObject{
		Geo: &proto.QPBxGeoObject_Point{
			Point: &proto.QPBxGeoPoint{
				Gid:        c.Gid,
				Feature:    c.Feature.Pb(),
				Properties: c.Properties,
			},
		},
	}
}
