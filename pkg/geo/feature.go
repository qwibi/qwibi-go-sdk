package geo

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QFeature[T geometry.QGeometry] struct {
	Geometry   T      `json:"geometry"`
	Properties []byte `json:"properties"`
}

func NewFeature[T geometry.QGeometry](geometry T) *QFeature[T] {
	return &QFeature[T]{
		Geometry:   geometry,
		Properties: []byte{},
	}
}

func NewFeaturePb(in *proto.QPBxFeature) (*QFeature[geometry.QGeometry], error) {
	geometry, err := geometry.NewGeometryPb(in.Geometry)
	if err != nil {
		return nil, qlog.Error(err)
	}

	feature := NewFeature(geometry)
	return feature, nil
}

func (c *QFeature[T]) Pb() *proto.QPBxFeature {
	return &proto.QPBxFeature{
		Geometry:   c.Geometry.Pb(),
		Properties: c.Properties,
	}
}

//func (f *QFeature[T]) UnmarshalJSON(data []byte) error {
//	qlog.Infof("!!!!! data: %+s", data)
//
//	type Alias QFeature[T]
//	aux := &struct {
//		Geometry map[string]interface{} `json:"geometry"`
//		*Alias
//	}{
//		Alias: (*Alias)(f),
//	}
//
//	if err := json.Unmarshal(data, &aux); err != nil {
//		return err
//	}
//
//	qlog.Infof("??? aux => %+v", aux)
//	qlog.Infof("??? geometry => %+v", aux.Geometry)
//	qlog.Infof("??? alias => %+v", aux.Alias)
//
//	if aux.Geometry == nil {
//		return qlog.Error("Error decoding Geometry property type nil")
//	}
//
//	switch v := aux.Geometry["type"]; v {
//	case geometry.QPointGeometryType:
//		var point geometry.QPoint
//		err := mapstructure.Decode(aux.Geometry, &point)
//		if err != nil {
//			return qlog.Error(err)
//		}
//		f.Geometry = point
//	default:
//		return qlog.Errorf("Unsupported geometry type: %s", v)
//	}
//
//	return nil
//}
