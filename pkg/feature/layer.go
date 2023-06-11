package feature

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QLayerFeature struct {
	Geometry   *geometry.QLayer `json:"geometry"`
	Properties []byte           `json:"properties"`
}

func NewLayerFeature() *QLayerFeature {
	return &QLayerFeature{
		Geometry: geometry.NewLayer(),
	}
}

func NewLayerFeaturePb(pb *proto.QPBxLayerFeature) (*QLayerFeature, error) {
	if pb == nil {
		return nil, qlog.Error("Invalid parameter type nil")
	}

	geometry, err := geometry.NewLayerPb(pb.Geometry)
	if err != nil {
		return nil, qlog.Error(err)
	}

	feature := &QLayerFeature{
		Geometry:   geometry,
		Properties: pb.Properties,
	}

	return feature, feature.Valid()
}

func (c *QLayerFeature) GetType() string {
	return QLayerFeatureType
}

func (c *QLayerFeature) Valid() error {
	if c.Geometry == nil {
		return qlog.Error("Invalid parameter type nil")
	}

	return c.Geometry.Valid()
}

func (c *QLayerFeature) Pb() *proto.QPBxLayerFeature {
	return &proto.QPBxLayerFeature{
		Geometry:   c.Geometry.Pb(),
		Properties: c.Properties,
	}
}

//func (c *QLayerFeature) MarshalJSON() ([]byte, error) {
//	type Alias QLayerFeature
//	return json.Marshal(&struct {
//		Type       string           `json:"type"`
//		Geometry   *geometry.QLayer `json:"geometry"`
//		Properties []byte           `json:"properties"`
//		*Alias
//	}{
//		Type:       "Feature",
//		Geometry:   c.Geometry,
//		Properties: c.Properties,
//	})
//}

//	func (c *QLayerFeature) UnmarshalJSON(data []byte) error {
//		var v map[string]interface{}
//		if err := json.Unmarshal(data, &v); err != nil {
//			return qlog.Error(err)
//		}
//		c.Geometry = geometry.NewLayer()
//		c.Properties = []byte{}
//		qlog.Infof("***** %+v", v)
//		return nil
//	}
