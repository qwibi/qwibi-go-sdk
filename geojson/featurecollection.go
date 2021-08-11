package geojson

import (
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QFeatureCollection ...
type QFeatureCollection struct {
	Features []*QFeature
}

// Pb ...
func (c *QFeatureCollection) Pb() (*proto.QPBxFeatureCollection, error) {
	features := []*proto.QPBxFeature{}
	for _, feature := range c.Features {
		pbFeature, err := feature.Pb()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		features = append(features, pbFeature)
	}
	pb := &proto.QPBxFeatureCollection{
		Features: features}

	return pb, nil
}

// NewFeatureCollection ...
func NewFeatureCollection() *QFeatureCollection {
	features := []*QFeature{}
	fc := &QFeatureCollection{
		Features: features}
	return fc
}

// Add ...
func (c *QFeatureCollection) Add(feature *QFeature) {
	c.Features = append(c.Features, feature)
}
