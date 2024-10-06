package point

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QPointFeatureProperties struct {
	ImageURI *string `json:"image_uri"`
}

func NewPointProperties() *QPointFeatureProperties {
	return &QPointFeatureProperties{}
}

func NewPointFeaturePropertiesPb(pb *proto.QPBxPointFeature_Properties) *QPointFeatureProperties {
	properties := NewPointProperties()
	properties.ImageURI = pb.ImageUri
	return properties
}

func NewPointFeaturePropertiesByte(b []byte) (*QPointFeatureProperties, error) {
	pb := &proto.QPBxPointFeature_Properties{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return nil, qlog.Error(err)
	}
	return NewPointFeaturePropertiesPb(pb), nil
}

func (c *QPointFeatureProperties) pb() *proto.QPBxPointFeature_Properties {
	return &proto.QPBxPointFeature_Properties{
		ImageUri: c.ImageURI,
	}
}
