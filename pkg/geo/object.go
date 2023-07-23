package geo

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/internal/utils"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QGeoObject struct {
	Gid        string             `json:"gid"`
	ParentID   string             `json:"parent_id,omitempty"`
	Geometry   geometry.QGeometry `json:"geometry"`
	Properties []byte             `json:"properties"`
}

func NewGeoObject(geometry geometry.QGeometry, options ...Option) *QGeoObject {
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
		ParentID:   "",
		Geometry:   geometry,
		Properties: config.Properties,
	}
}

func NewGeoObjectPb(in *proto.QPBxGeoObject) (*QGeoObject, error) {
	object := &QGeoObject{
		Gid:        in.Gid,
		ParentID:   in.ParentID,
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
		ParentID:   c.ParentID,
		Properties: c.Properties,
	}
}

func NewGeoObjectBytes(data []byte) (*QGeoObject, error) {
	var raw QGeoObject

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, qlog.Error(err)
	}

	qlog.TODO()

	return &raw, nil
}

//func (c *QGeoObject) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		Gid string `json:"type"`
//	}{
//		Gid: utils.NewID(),
//	})
//}

func (c *QGeoObject) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
