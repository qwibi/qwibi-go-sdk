package feature

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
)

type QLayerFeature struct {
	Properties []byte
}

func NewLayerFeature() *QLayerFeature {
	return &QLayerFeature{
		Properties: nil,
	}
}

func NewLayerFeaturePb(pb *proto.QPBxLayerFeature) (*QLayerFeature, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	feature := &QLayerFeature{}

	return feature, feature.Valid()
}

func (c *QLayerFeature) Valid() error {
	return qlog.TODO()
}

func (c *QLayerFeature) Pb() (*proto.QPBxLayerFeature, error) {

	featurePb := &proto.QPBxLayerFeature{
		Properties: c.Properties,
	}

	return featurePb, nil
}
