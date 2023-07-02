package geometry

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QLayer ...
type QLayer struct{}

// NewLayer ...
func NewLayer() *QLayer {
	return &QLayer{}
}

// NewLayerPb ...
func NewLayerPb(pb *proto.QPBxLayer) (*QLayer, error) {
	if pb == nil {
		err := errors.New("Invalid parameter type nil")
		return nil, qlog.Error(err)
	}

	layer := NewLayer()

	return layer, layer.Valid()
}

// Valid ...
func (c *QLayer) GetType() string {
	return QLayerGeometryType
}

// Valid ...
func (c *QLayer) Valid() error {
	return nil
}

// Pb ...
func (c *QLayer) Pb() *proto.QPBxGeometry {
	return &proto.QPBxGeometry{
		Type: &proto.QPBxGeometry_Layer{
			Layer: &proto.QPBxLayer{},
		},
	}
}

//func (f *QLayer) UnmarshalJSON(data []byte) error {
//	qlog.TODO("#### Layer")
//	return nil
//}
