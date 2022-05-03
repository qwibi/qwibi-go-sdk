package geometry

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

// QLine ...
type QLayer struct {
	Properties *structpb.Struct
}

// NewLayer
func NewDefaultLayer() *QLayer {
	return &QLayer{
		Properties: &structpb.Struct{},
	}
}

func NewLayer(properties *structpb.Struct) (*QLayer, error) {
	return &QLayer{
		Properties: properties,
	}, nil
}

// NewLayerPb ...
func NewLayerPb(pb *proto.QPBxLayer) (*QLayer, error) {
	if pb == nil {
		err := errors.New("Invalid layer format type nil")
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}
	return NewLayer(pb.Properties)
}

// Valid ...
func (c *QLayer) Valid() error {
	return nil
}

// Pb ...
func (c *QLayer) Pb() (*proto.QPBxGeometry, error) {

	pb := &proto.QPBxGeometry{
		Geometry: &proto.QPBxGeometry_Layer{
			Layer: &proto.QPBxLayer{
				Properties: c.Properties,
			},
		},
	}

	return pb, nil
}
