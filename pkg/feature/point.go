package feature

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QPointFeature struct {
	Geometry   *geometry.QPoint `json:"geometry"`
	Properties []byte           `json:"properties"`
}

func NewPointFeature() *QPointFeature {
	return &QPointFeature{
		Geometry: geometry.NewPoint(),
	}
}

func NewPointFeaturePb(pb *proto.QPBxPointFeature) (*QPointFeature, error) {
	if pb == nil {
		return nil, qlog.Error("Invalid parameter type nil")
	}

	geometry, err := geometry.NewPointPb(pb.Geometry)
	if err != nil {
		return nil, qlog.Error(err)
	}

	feature := &QPointFeature{
		Geometry:   geometry,
		Properties: pb.Properties,
	}

	return feature, feature.Valid()
}

func (c *QPointFeature) GetType() string {
	return QPointFeatureType
}

func (c *QPointFeature) Valid() error {
	if c.Geometry == nil {
		return qlog.Error("Invalid parameter type nil")
	}

	return c.Geometry.Valid()
}

func (c *QPointFeature) Pb() *proto.QPBxPointFeature {
	return &proto.QPBxPointFeature{
		Geometry:   c.Geometry.Pb(),
		Properties: c.Properties,
	}
}

//func (c *QPointFeature) MarshalJSON() ([]byte, error) {
//	type Alias QPointFeature
//	return json.Marshal(&struct {
//		Type       string           `json:"type"`
//		Geometry   *geometry.QPoint `json:"geometry"`
//		Properties []byte           `json:"properties"`
//		*Alias
//	}{
//		Type:       "Feature",
//		Geometry:   c.Geometry,
//		Properties: c.Properties,
//	})
//}

//	func (c *QPointFeature) UnmarshalJSON(data []byte) error {
//		var v map[string]interface{}
//		if err := json.Unmarshal(data, &v); err != nil {
//			return qlog.Error(err)
//		}
//		c.Geometry = geometry.NewPoint()
//		c.Properties = []byte{}
//		qlog.Infof("***** %+v", v)
//		return nil
//	}
