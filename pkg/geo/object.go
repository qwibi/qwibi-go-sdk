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

func NewGeoObject(features []*feature.QFeature, option ...ObjectOption) (*QGeoObject, error) {
	if features == nil {
		return nil, qlog.Error("features is not defined")
	}

	h := &QGeoObject{
		Features: features,
	}

	for _, opt := range option {
		opt(h)
	}

	return h, nil
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

	return NewGeoObject(features,
		WithGid(in.Gid),
		WithProperties(in.Properties),
	)
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
