package feature

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QFeature interface {
	GetType() string
	//SetCenter(float64, float64)
	//Pb() *proto.QPBxFeature
	//UnmarshalJSON(data []byte) error
	//MarshalJSON() ([]byte, error)
}

func NewFeaturePb(pb *proto.QPBxFeature) (QFeature, error) {
	switch v := pb.Type.(type) {
	case *proto.QPBxFeature_Point:
		return NewPointFeaturePb(v.Point)
	default:
		return nil, qlog.Errorf("Unknown Feature type: %v", v)
	}
}

//func NewFeatureDecoder(i interface{}) (QFeature, error) {
//	switch v := i.(type) {
//	case map[string]interface{}:
//		if p1, ok := v["type"]; ok {
//			if p1.(string) == "Feature" {
//				if p2, ok := v["geometry"]; ok {
//					p2
//				}
//			}
//		}
//
//	default:
//		return nil, qlog.Error("Unknown Feature type: %s", v)
//	}
//
//	return nil, nil
//}
