package geo

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/feature"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type QGeoObject struct {
	Gid        string
	Features   []*feature.QFeature
	Properties map[string]interface{}
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
		WithProperties(in.Properties.AsMap()),
	)
}

func (c *QGeoObject) Pb() *proto.QPBxGeoObject {
	featuresPb := []*proto.QPBxFeature{}
	for _, feature := range c.Features {
		featuresPb = append(featuresPb, feature.FeaturePb())
	}

	pbProperties, err := structpb.NewStruct(c.Properties)
	if err != nil {
		qlog.Error(err)
	}

	return &proto.QPBxGeoObject{
		Gid:        c.Gid,
		Features:   featuresPb,
		Properties: pbProperties,
	}
}
