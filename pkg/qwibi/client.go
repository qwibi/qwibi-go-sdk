package qwibi

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	grpc "google.golang.org/grpc"
)

type QApiClient struct {
	apiClient proto.QPBxApiClient
	ctx       context.Context
}

func NewClient(ctx context.Context, addr string) (*QApiClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		conn.Close()
		return nil, errors.Wrap(err, "Connection error")
	}

	client := &QApiClient{
		apiClient: proto.NewQPBxApiClient(conn),
		ctx:       ctx,
	}

	return client, nil
}

func (c *QApiClient) Layer(options ...layer.LayerOption) (*geoLayer, error) {
	g := layer.NewGeoLayer(options...)

	cmdListPb := []*proto.QPBxCommand{}
	for _, cmd := range g.Commands() {
		cmdListPb = append(cmdListPb, cmd.Pb())
	}

	req := &proto.QPBxLayerRequest{
		LayerId:    g.LayerId(),
		Public:     g.Public(),
		Properties: g.Properties(),
		Commands:   cmdListPb,
	}

	res, err := c.apiClient.Layer(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	layer, err := layer.NewGeoLayerPb(res.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	l := &geoLayer{
		layer:      layer,
		QApiClient: c,
	}

	return l, nil
}
