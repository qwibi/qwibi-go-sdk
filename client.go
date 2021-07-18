package sdk

import (
	"context"
	"io"
	"runtime"

	"github.com/rs/zerolog/log"
	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/auth"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/rpc/request"
	"github.com/qwibi/qwibi-go-sdk/rpc/response"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type QApiClient struct {
	apiClient proto.QPBxApiClient
	ctx       context.Context
}

func NewClient(addr string) (*QApiClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		conn.Close()
		return nil, errors.Wrap(err, "Connection error")
	}

	client := &QApiClient{
		apiClient: proto.NewQPBxApiClient(conn),
		ctx:       context.Background(),
	}

	return client, nil
}

// Auth ...
func (c *QApiClient) Auth(req *request.QAuthRequest) (*response.QAuthResponse, error) {

	reqPb, err := req.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resPb, err := c.apiClient.Auth(c.ctx, reqPb)
	if err != nil {
		return nil, errors.Wrap(err, "Auth error")
	}

	res, err := response.NewAuthResponsePb(resPb)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	md := metadata.Pairs("token", res.Token)
	c.ctx = metadata.NewOutgoingContext(context.Background(), md)

	return res, nil
}

func (c *QApiClient) BasicAuth(login string, password string) (*response.QAuthResponse, error) {
	auth, err := auth.NewBasicAuth(login, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req := &request.QAuthRequest{
		Auth: auth,
	}

	return c.Auth(req)
}

// Join ...
func (c *QApiClient) Join(layerID string) (*geo.QLayer, error) {
	joinRequest, err := request.NewJoinRequest(layerID)
	if err != nil {
		return nil, err
	}

	requestPb, err := joinRequest.Pb()
	if err != nil {
		return nil, err
	}

	joinResponse, err := c.apiClient.Join(c.ctx, requestPb)
	if err != nil {
		return nil, err
	}

	layer, err := geo.NewLayerPb(joinResponse.Layer)
	if err != nil {
		return nil, err
	}

	return layer, nil
}

// Stream
func (c *QApiClient) Stream(ctx context.Context ) error {
	streamRequest, err := request.NewStreamRequest()
	if err != nil {
		return errors.WithStack(err)
	}

	pb, err := streamRequest.Pb()
	if err != nil {
		return errors.WithStack(err)
	}

	stream, err := c.apiClient.Stream(ctx, pb)
	if err != nil {
		return errors.WithStack(err)
	}

	//go c.send(stream)
	go c.receive(stream)
	//return c.RunTest(ctx)
	//go c.RunTest(ctx)

	return nil
}

// Receive ...
func (c *QApiClient) receive(stream proto.QPBxApi_StreamClient) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Info().Msgf("Client connected: %+v", c)
			log.Info().Msgf("Goroutines: %d", runtime.NumGoroutine())
			return err
		}
		if err != nil {
			errors.WithStack(err)
		}

		log.Info().Msgf("==> Stream: [%T] %+v", msg.Response, msg)
	}
}