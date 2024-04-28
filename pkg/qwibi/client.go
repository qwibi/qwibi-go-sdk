package qwibi

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/metadata"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/rpc"
	"github.com/qwibi/qwibi-go-sdk/pkg/rpc/request"
	"github.com/qwibi/qwibi-go-sdk/pkg/rpc/response"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	grpc "google.golang.org/grpc"
	"io"
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

// Bot ...
func (c *QApiClient) Bot(token string) (*geoBot, error) {
	c.ctx = metadata.SetBotContextToken(c.ctx, token)

	client, err := c.apiClient.Bot(c.ctx, grpc.EmptyCallOption{})
	if err != nil {
		return nil, qlog.Error(err)
	}

	bot := &geoBot{
		client: client,
	}

	return bot, nil
}

func (c *QApiClient) Token(layerId string) (utils.Token, error) {
	req := &proto.QPBxTokenRequest{
		LayerId: layerId,
	}

	res, err := c.apiClient.Token(c.ctx, req)
	if err != nil {
		return "", qlog.Error(err)
	}

	return res.Token, nil
}

func (c *QApiClient) Layer(options ...layer.LayerOption) (*response.QLayerResponse, error) {

	req, err := request.NewLayerRequest(options...)
	if err != nil {
		return nil, qlog.Error(err)
	}

	resPb, err := c.apiClient.Layer(c.ctx, req.Pb())
	if err != nil {
		return nil, qlog.Error(err)
	}

	layer, err := layer.NewGeoLayerPb(resPb.Layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	res, err := response.NewLayerResponse(req.RequestId, layer)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return res, nil
}

func (c *QApiClient) Post(layerId string, object *geo.QGeoObject) (*response.QPostResponse, error) {

	req, err := request.NewPostRequest(layerId, object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	resPb, err := c.apiClient.Post(c.ctx, req.Pb())
	if err != nil {
		return nil, qlog.Error(err)
	}

	res, err := response.NewPostResponsePb(resPb)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return res, nil
}

// Subscribe ...
func (c *QApiClient) Subscribe(handler func(messge rpc.QEvent)) error {
	if c == nil {
		qlog.Error("bad layer parameter type nil")
	}

	if c == nil {
		qlog.Error("bad layer parameter type nil")
	}

	client, err := c.apiClient.Stream(c.ctx, grpc.EmptyCallOption{})
	if err != nil {
		return qlog.Error(err)
	}

	for {
		in, err := client.Recv()
		if err == io.EOF {
			//qlog.Infof("Client connected: %+v", c)
			//qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
			return qlog.Error(err)
		}
		if err != nil {
			return qlog.Error(err)
		}

		//qlog.Debugf("Client stream: [%T] %v", in, in)
		message, err := rpc.NewEventPb(in.Event)
		if err != nil {
			return qlog.Error(err)
		}

		handler(message)
	}
}
