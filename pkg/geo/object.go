package geo

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/feature"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QGeoObject struct {
	Gid        string
	Features   []*feature.QFeature
	Properties []byte
}

func NewGeoObject(gid string, features []*feature.QFeature, properties []byte) (*QGeoObject, error) {
	if features == nil {
		return nil, qlog.Error("features is not defined")
	}

	object := &QGeoObject{
		Gid:        gid,
		Features:   features,
		Properties: properties,
	}

	return object, nil
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
	features := make([]*feature.QFeature, 0)

	for _, pb := range in.Features {
		feature, err := feature.NewFeaturePb(pb)
		if err != nil {
			return nil, qlog.Error(err)
		}
		features = append(features, feature)
	}

	return NewGeoObject(in.Gid, features, in.Properties)
}

func (c *QGeoObject) Pb() *proto.QPBxGeoObject {
	featuresPb := []*proto.QPBxFeature{}
	for _, feature := range c.Features {
		featuresPb = append(featuresPb, feature.FeaturePb())
	}
	return &proto.QPBxGeoObject{
		Gid:        c.Gid,
		Features:   featuresPb,
		Properties: c.Properties,
	}
}
