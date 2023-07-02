package geo

//
//import (
//	"encoding/json"
//	"github.com/qwibi/qwibi-go-sdk/internal/utils"
//	"github.com/qwibi/qwibi-go-sdk/pkg/feature"
//	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
//	"github.com/qwibi/qwibi-go-sdk/proto"
//	"regexp"
//)
//
//type QGeoPoint struct {
//	Gid        string                 `json:"gid"`
//	Feature    *feature.QPointFeature `json:"feature"`
//	Properties []byte                 `json:"properties"`
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
////func NewGeoPoint(feature feature.QPointFeature) (*QGeoObject, error) {
////	object := &QGeoObject{
////		Gid:        utils.NewID(),
////		Feature:    feature,
////		Properties: []byte{},
////	}
////
////	return object, object.Valid()
////}
//
//func NewGeoPoint() *QGeoPoint {
//	return &QGeoPoint{
//		Gid:     utils.NewID(),
//		Feature: feature.NewPointFeature(),
//	}
//}
//
//func NewGeoPointData(data []byte) (*QGeoPoint, error) {
//	object := &QGeoPoint{}
//	err := json.Unmarshal(data, &object)
//	if err != nil {
//		return nil, err
//	}
//
//	return object, nil
//}
//
////func NewGeoPointPb(in *proto.QPBxGeoPoint) (*QGeoPoint, error) {
////	feature, err := feature.NewPointFeaturePb(in.Feature)
////	if err != nil {
////		return nil, qlog.Error(err)
////	}
////
////	if in.Feature == nil {
////		return nil, qlog.Error("Feature not defined")
////	}
////
////	object := &QGeoPoint{
////		Gid:        in.Gid,
////		Feature:    feature,
////		Properties: in.Properties,
////	}
////
////	return object, object.Valid()
////}
//
//func (c *QGeoPoint) GetType() string {
//	return QGeoPointType
//}
//
//func (c *QGeoPoint) GetGid() string {
//	return c.Gid
//}
//
//func (c *QGeoPoint) Pb() *proto.QPBxGeoObject {
//	return &proto.QPBxGeoObject{
//		Gid:        c.Gid,
//		Feature:    c.Feature.Pb(),
//		Properties: c.Properties,
//	}
//}
//
//func (c *QGeoPoint) Valid() error {
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
//func (c *QGeoPoint) MarshalJSON() ([]byte, error) {
//	data := map[string]interface{}{
//		"type":       c.GetType(),
//		"gid":        c.Gid,
//		"feature":    c.Feature,
//		"properties": c.Properties,
//	}
//	return json.Marshal(data)
//}
