package feature

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/feature/point"
	sdkGeometry "github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// Feature ...
type Feature interface {
	Fid() string
	Geometry() *sdkGeometry.QGeometry
	FeaturePb() *proto.QPBxFeature
	Properties() []byte
	Valid() error
}

type QFeature struct {
	Feature
}

func NewFeature(fid string, geometry *sdkGeometry.QGeometry, properties []byte) *QFeature {
	switch v := geometry.Geometry.(type) {
	case *sdkGeometry.QPoint:
		pointFeature, err := point.NewPointFeature(v,
			point.WithFid(fid),
			point.WithProperties(properties),
		)
		if err != nil {
			qlog.Error(err)
			return nil
		}
		return &QFeature{pointFeature}
	default:
		qlog.Warnf("unknown geometry type: %T", v)
	}

	return nil
}

// NewGeometryPb ...
func NewFeaturePb(in *proto.QPBxFeature) (*QFeature, error) {
	if in == nil {
		return nil, qlog.Errorf("invalid pointFeature type nil")
	}

	switch v := in.Type.(type) {
	case *proto.QPBxFeature_Point:
		pointFeature, err := point.NewPointFeaturePb(v.Point)
		if err != nil {
			return nil, qlog.Error(err)
		}
		return &QFeature{pointFeature}, nil
	default:
		return nil, qlog.Errorf("unknown pointFeature type: %s", v)
	}
}

//func (c *QFeature) String() string {
//	return c.FeaturePb().String()
//}
