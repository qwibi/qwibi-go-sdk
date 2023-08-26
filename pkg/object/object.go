package object

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QGeoObject struct {
	Gid        string              `json:"gid"`
	Geometry   *geometry.QGeometry `json:"geometry"`
	Properties []byte              `json:"properties"`
}

func NewGeoObject(options ...Option) *QGeoObject {
	// Default configuration
	config := &Config{
		Gid: utils.NewID(),
	}

	// Apply options
	for _, opt := range options {
		opt(config)
	}

	return &QGeoObject{
		Gid:        config.Gid,
		Geometry:   config.Geometry,
		Properties: config.Properties,
	}
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
	object := &QGeoObject{
		Gid:        in.Gid,
		Properties: in.Properties,
	}

	geometry, err := geometry.NewGeometryPb(in.Geometry)
	if err != nil {
		return nil, qlog.Error(err)
	}

	object.Geometry = geometry

	return object, nil
}

func (c *QGeoObject) Pb() *proto.QPBxGeoObject {
	return &proto.QPBxGeoObject{
		Gid:        c.Gid,
		Geometry:   c.Geometry.Pb(),
		Properties: c.Properties,
	}
}

func NewGeoObjectBytes(data []byte) (*QGeoObject, error) {
	var object QGeoObject

	err := json.Unmarshal(data, &object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return &object, nil
}

func (c *QGeoObject) Valid() error {
	if c.Geometry == nil {
		return qlog.Error("Object error: geometry not defined")
	}

	return nil
}

func (c *QGeoObject) UnmarshalJSON(data []byte) error {

	var raw map[string]interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c.Gid, _ = raw["gid"].(string)
	c.Properties = []byte(raw["properties"].(string))
	rawGeometry := raw["geometry"].(map[string]interface{})
	geometry, err := geometry.NewGeometryStruct(rawGeometry)
	if err != nil {
		return err
	}
	c.Geometry = geometry

	return nil
}

//func (c *QGeoObject) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		Gid        string              `json:"gid"`
//		Geometry   *geometry.QGeometry `json:"geometry"`
//		Properties string              `json:"properties"`
//	}{
//		Gid:        c.Gid,
//		Geometry:   c.Geometry,
//		Properties: string(c.Properties),
//	})
//}

//func (c *QGeoObject) String() string {
//	b, _ := json.Marshal(c)
//	return string(b)
//}
