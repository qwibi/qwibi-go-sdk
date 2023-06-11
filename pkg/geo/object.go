package geo

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QGeoObject interface {
	GetGid() string
	GetType() string
	Pb() *proto.QPBxGeoObject
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (QGeoObject, error) {
	switch v := in.Type.(type) {
	case *proto.QPBxGeoObject_Layer:
		return NewGeoLayerPb(v.Layer)
	case *proto.QPBxGeoObject_Point:
		return NewGeoPointPb(v.Point)
	default:
		return nil, qlog.Errorf("Unknown Feature type: %v", v)
	}
}

func NewGeoObjectData(in []byte) (QGeoObject, error) {

	var raw map[string]interface{}
	err := json.Unmarshal(in, &raw)
	if err != nil {
		return nil, qlog.Error(err)
	}

	t, ok := raw["type"].(string)
	if !ok {
		return nil, qlog.Error("Object type not defined")
	}

	switch t {
	case QGeoLayerType:
		return NewGeoLayerData(in)
	case QGeoPointType:
		return NewGeoPointData(in)
	default:
		return nil, qlog.Error("Unknown object type...", t)
	}
}

//package geo
//
//import (
//	"github.com/qwibi/qwibi-go-sdk/internal/utils"
//	"github.com/qwibi/qwibi-go-sdk/pkg/feature"
//	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
//	"github.com/qwibi/qwibi-go-sdk/proto"
//	regexp "github.com/wasilibs/go-re2"
//)
//
//type QGeoObject struct {
//	Gid        string           `json:"gid"`
//	Feature    feature.QFeature `json:"feature"`
//	Properties []byte           `json:"properties"`
//}
//
////func (c *QGeoObject) UnmarshalJSON(data []byte) error {
////	var v map[string]interface{}
////	if err := json.Unmarshal(data, &v); err != nil {
////		return qlog.Error(err)
////	}
////
////	if p1, ok := v["gid"]; ok {
////		c.Gid = p1.(string)
////	}
////
////	if p2, ok := v["feature"]; ok {
////		feature, err := feature2.NewFeatureDecoder(p2)
////		if err != nil {
////			return qlog.Error(err)
////		}
////		c.Feature = feature
////	}
////
////	qlog.Infof("!!!! %+s", v)
////
////	return nil
////}
//
//func NewGeoObject(feature feature.QFeature) (*QGeoObject, error) {
//	object := &QGeoObject{
//		Gid:        utils.NewID(),
//		Feature:    feature,
//		Properties: []byte{},
//	}
//
//	return object, object.Valid()
//}
//
//func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
//	object := &QGeoObject{
//		Gid:        in.Gid,
//		Properties: in.Properties,
//	}
//
//	if in.Feature == nil {
//		return nil, qlog.Error("Feature not defined")
//	}
//
//	if in.Feature.Type == nil {
//		return nil, qlog.Error("Invalid Feature format")
//	}
//
//	switch v := in.Feature.Type.(type) {
//	case *proto.QPBxFeature_Point:
//		f, err := feature.NewPointFeaturePb(v.Point)
//		if err != nil {
//			return nil, qlog.Error(err)
//		}
//		object.Feature = f
//	default:
//		qlog.Error("Unknown feature type", v)
//	}
//
//	return object, object.Valid()
//}
//
//func (c *QGeoObject) Pb() *proto.QPBxGeoObject {
//	return &proto.QPBxGeoObject{
//		Gid:        c.Gid,
//		Feature:    c.Feature.Pb(),
//		Properties: c.Properties,
//	}
//}
//
//func (c *QGeoObject) Valid() error {
//	if c.Gid == "" {
//		return qlog.Error("Object gid not defined")
//	}
//
//	if c.Feature == nil {
//		return qlog.Error("Feature not defined")
//	}
//
//	re := regexp.MustCompile("^[A-Za-z0-9_]\\w{4,20}$")
//	if !re.MatchString(c.Gid) {
//		return qlog.Error("Wrong gid format:", c.Gid)
//	}
//
//	return nil
//}
//
//func NewGeoPoint() *QGeoObject {
//	object, _ := NewGeoObject(feature.NewPointFeature())
//	return object
//}
