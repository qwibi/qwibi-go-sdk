package geo

import (
	"encoding/json"
	"github.com/qwibi/qwibi-go-sdk/internal/utils"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"regexp"
)

//import (
//	"encoding/json"
//	"github.com/qwibi/qwibi-go-sdk/internal/utils"
//	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
//	"github.com/qwibi/qwibi-go-sdk/proto"
//	"regexp"
//)
//

type QGeoLayer struct {
	Gid        string        `json:"gid"`
	Objects    []*QGeoObject `json:"objects"`
	Properties []byte        `json:"properties"`
}

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

func NewGeoLayerPb(in *proto.QPBxGeoLayer) (*QGeoLayer, error) {
	var objects []*QGeoObject

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

func NewGeoLayerBytes(data []byte) (*QGeoLayer, error) {
	var layer QGeoLayer
	err := json.Unmarshal(data, &layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return &layer, nil
}

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
