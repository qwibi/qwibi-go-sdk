package geo

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/internal/utils"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"regexp"
)

type QGeoLayer struct {
	Gid        string                            `json:"gid"`
	Objects    []*QGeoObject[geometry.QGeometry] `json:"objects"`
	Properties []byte                            `json:"properties"`
}

//func (c *QGeoObject) UnmarshalJSON(data []byte) error {
//	var v map[string]interface{}
//	if err := json.Unmarshal(data, &v); err != nil {
//		return qlog.Error(err)
//	}
//
//	if p1, ok := v["gid"]; ok {
//		c.Gid = p1.(string)
//	}
//
//	if p2, ok := v["feature"]; ok {
//		feature, err := feature2.NewFeatureDecoder(p2)
//		if err != nil {
//			return qlog.Error(err)
//		}
//		c.Feature = feature
//	}
//
//	qlog.Infof("!!!! %+s", v)
//
//	return nil
//}

//func NewGeoLayer(feature feature.QLayerFeature) (*QGeoObject, error) {
//	object := &QGeoObject{
//		Gid:        utils.NewID(),
//		Feature:    feature,
//		Properties: []byte{},
//	}
//
//	return object, object.Valid()
//}

func NewGeoLayer(options ...Option) *QGeoLayer {
	// Default configuration
	config := &Config{
		Gid: utils.NewID(),
	}

	// Apply options
	for _, opt := range options {
		opt(config)
	}

	return &QGeoLayer{
		Gid: config.Gid,
	}
}

func NewGeoLayerData(data []byte) (*QGeoLayer, error) {
	object := &QGeoLayer{}
	err := json.Unmarshal(data, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func NewGeoLayerPb(in *proto.QPBxGeoLayer) (*QGeoLayer, error) {
	var objects []*QGeoObject[geometry.QGeometry]

	for _, k := range in.Objects {
		object, err := NewGeoObjectPb(k)
		if err != nil {
			return nil, qlog.Error(err)
		}

		objects = append(objects, object)
	}

	object := &QGeoLayer{
		Gid:        in.Gid,
		Objects:    objects,
		Properties: in.Properties,
	}

	return object, object.Valid()
}

//func (c *QGeoLayer) GetType() string {
//	return QGeoLayerType
//}
//
//func (c *QGeoLayer) GetGid() string {
//	return c.Gid
//}

func (c *QGeoLayer) Pb() *proto.QPBxGeoLayer {
	var objectsPb []*proto.QPBxGeoObject
	for _, k := range c.Objects {
		objectPb := k.Pb()
		objectsPb = append(objectsPb, objectPb)
	}

	layerPb := &proto.QPBxGeoLayer{
		Gid:        c.Gid,
		Objects:    objectsPb,
		Properties: c.Properties,
	}

	return layerPb
}

func (c *QGeoLayer) Valid() error {
	if c.Gid == "" {
		return qlog.Error("Object gid not defined")
	}

	re := regexp.MustCompile("^[A-Za-z0-9_]\\w{3,20}$")
	if !re.MatchString(c.Gid) {
		return qlog.Error("Wrong gid format:", c.Gid)
	}

	return nil
}

func (c *QGeoLayer) Post(object *QGeoObject[geometry.QGeometry]) error {

	return qlog.TODO()
}

//func (c QGeoLayer) MarshalJSON() ([]byte, error) {
//	data := map[string]interface{}{
//		"type":       c.GetType(),
//		"gid":        c.Gid,
//		"feature":    c.Feature,
//		"properties": c.Properties,
//	}
//	return json.Marshal(data)
//}
