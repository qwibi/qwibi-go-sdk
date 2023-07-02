package geo

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/internal/utils"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"regexp"
)

type QGeoObject[T geometry.QGeometry] struct {
	Gid        string       `json:"gid"`
	Feature    *QFeature[T] `json:"feature"`
	Properties []byte       `json:"properties"`
}

//func NewGeoObject2[T feature.QFeature]() *QGeoObject {
//	return &QGeoObject{
//		Gid:        utils.NewID(),
//		Feature:    make(T),
//		Properties: []byte{},
//	}
//}

func NewGeoObject[T geometry.QGeometry](geometry T) *QGeoObject[T] {
	return &QGeoObject[T]{
		Gid:        utils.NewID(),
		Feature:    NewFeature(geometry),
		Properties: []byte{},
	}
}

func NewGeoPoint() *QGeoObject[geometry.QPoint] {
	return NewGeoObject(geometry.NewPoint())
}

//func NewGeoObject(feature *QFeature) *QGeoObject {
//	return &QGeoObject{
//		Gid:        utils.NewID(),
//		Feature:    feature,
//		Properties: []byte{},
//	}
//}

//func NewGeoPoint() *QGeoObject {
//	return NewGeoObject(feature.NewPointFeature())
//}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject[geometry.QGeometry], error) {
	feature, err := NewFeaturePb(in.Feature)
	if err != nil {
		return nil, qlog.Error(err)
	}

	object := &QGeoObject[geometry.QGeometry]{
		Gid:        in.Gid,
		Feature:    feature,
		Properties: in.Properties,
	}

	return object, object.Valid()
}

func NewGeoObjectBytes(data []byte) (*QGeoObject[geometry.QGeometry], error) {

	var object QGeoObject[geometry.QGeometry]
	var raw QGeoObject[geometry.QGeometry]
	//var raw map[string]interface{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, qlog.Error(err)
	}

	qlog.Debugf("NewGeoObjectBytes (Feature): %+v", raw.Feature)

	//t, ok := raw["type"].(string)
	//if !ok {
	//	return nil, qlog.Error("Object type not defined")
	//}
	//
	//switch t {
	//case QGeoLayerType:
	//	return NewGeoLayerData(data)
	//case QGeoPointType:
	//	return NewGeoPointData(data)
	//default:
	//	return nil, qlog.Error("Unknown object type...", t)
	//}

	return &object, qlog.TODO("check me")
}

// package geo
//
// import (
//
//	"github.com/qwibi/qwibi-go-sdk/internal/utils"
//	"github.com/qwibi/qwibi-go-sdk/pkg/feature"
//	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
//	"github.com/qwibi/qwibi-go-sdk/proto"
//	regexp "github.com/wasilibs/go-re2"
//
// )
//
//	type QGeoObject struct {
//		Gid        string           `json:"gid"`
//		Feature    feature.QFeature `json:"feature"`
//		Properties []byte           `json:"properties"`
//	}
//
// //func (c *QGeoObject) UnmarshalJSON(data []byte) error {
// //	var v map[string]interface{}
// //	if err := json.Unmarshal(data, &v); err != nil {
// //		return qlog.Error(err)
// //	}
// //
// //	if p1, ok := v["gid"]; ok {
// //		c.Gid = p1.(string)
// //	}
// //
// //	if p2, ok := v["feature"]; ok {
// //		feature, err := feature2.NewFeatureDecoder(p2)
// //		if err != nil {
// //			return qlog.Error(err)
// //		}
// //		c.Feature = feature
// //	}
// //
// //	qlog.Infof("!!!! %+s", v)
// //
// //	return nil
// //}
//
//	func NewGeoObject(feature feature.QFeature) (*QGeoObject, error) {
//		object := &QGeoObject{
//			Gid:        utils.NewID(),
//			Feature:    feature,
//			Properties: []byte{},
//		}
//
//		return object, object.Valid()
//	}
//
//	func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
//		object := &QGeoObject{
//			Gid:        in.Gid,
//			Properties: in.Properties,
//		}
//
//		if in.Feature == nil {
//			return nil, qlog.Error("Feature not defined")
//		}
//
//		if in.Feature.Type == nil {
//			return nil, qlog.Error("Invalid Feature format")
//		}
//
//		switch v := in.Feature.Type.(type) {
//		case *proto.QPBxFeature_Point:
//			f, err := feature.NewPointFeaturePb(v.Point)
//			if err != nil {
//				return nil, qlog.Error(err)
//			}
//			object.Feature = f
//		default:
//			qlog.Error("Unknown feature type", v)
//		}
//
//		return object, object.Valid()
//	}

func (c *QGeoObject[T]) Pb() *proto.QPBxGeoObject {
	return &proto.QPBxGeoObject{
		Gid:        c.Gid,
		Feature:    c.Feature.Pb(),
		Properties: c.Properties,
	}
}

func (c *QGeoObject[T]) Valid() error {
	if c.Gid == "" {
		return qlog.Error("Object gid not defined")
	}

	if c.Feature == nil {
		return qlog.Error("Feature not defined")
	}

	re := regexp.MustCompile("^[A-Za-z0-9_]\\w{4,20}$")
	if !re.MatchString(c.Gid) {
		return qlog.Error("Wrong gid format:", c.Gid)
	}

	return nil
}

//
//func NewGeoPoint() *QGeoObject {
//	object, _ := NewGeoObject(feature.NewPointFeature())
//	return object
//}

//func (c *QGeoObject) UnmarshalJSON(data []byte) error {
//	v := struct {
//		Gid        string          `json:"gid"`
//		Feature    json.RawMessage `json:"feature"`
//		Properties []byte          `json:"properties"`
//		//*QGeoObject
//	}{}
//	if err := json.Unmarshal(data, &v); err != nil {
//		return qlog.Error(err)
//	}
//
//	//if p1, ok := v["gid"]; ok {
//	//	c.Gid = p1.(string)
//	//}
//	//
//	//if p2, ok := v["feature"]; ok {
//	//	feature, err := feature2.NewFeatureDecoder(p2)
//	//	if err != nil {
//	//		return qlog.Error(err)
//	//	}
//	//	c.Feature = feature
//	//}
//
//	qlog.Infof("!!!! %+s", v)
//
//	return nil
//}
