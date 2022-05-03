package geo

import (
	"github.com/qwibi/qwibi-go-sdk/geometry"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"google.golang.org/protobuf/types/known/structpb"
)

// NewGeoPoint ...
func NewGeoPoint() *QGeoObject {
	return &QGeoObject{
		Gid:        utils.NewID(),
		Geometry:   geometry.NewZeroPoint(),
		Properties: &structpb.Struct{},
	}
}
