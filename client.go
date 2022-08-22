package sdk

import (
	"context"
	"fmt"
	"io"
	"runtime"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/auth"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/rpc/request"
	"github.com/qwibi/qwibi-go-sdk/rpc/response"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"github.com/rs/zerolog/log"
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
		return nil, errors.WithStack(err)
	}

	res, err := response.NewAuthResponsePb(resPb)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	md := metadata.Pairs("token", res.Session.Token)
	c.ctx = metadata.NewOutgoingContext(context.Background(), md)

	return res, nil
}

func (c *QApiClient) AnonymousAuth() (*response.QAuthResponse, error) {
	auth, err := auth.NewAnonymousAuth()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req := &request.QAuthRequest{
		Auth: auth,
	}

	return c.Auth(req)
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
func (c *QApiClient) Join(gid string) (geo.QGeoObject, error) {
	if gid == "" {
		gid = utils.NewID()
	}

	joinRequest, err := request.NewJoinRequest(gid)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	requestPb, err := joinRequest.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	joinResponse, err := c.apiClient.Join(c.ctx, requestPb)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	object, err := geo.NewGeoObjectPb(joinResponse.Object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return object, nil
}

// Join ...
func (c *QApiClient) Post(object geo.QGeoObject) (geo.QGeoObject, error) {
	request, err := request.NewPostRequest(object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Debug().Msgf("Post request: %#v", request)
	responsePb, err := c.apiClient.Post(c.ctx, request.Pb())
	if err != nil {
		err := errors.New(fmt.Sprint(err))
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	response, err := response.NewPostResponsePb(responsePb)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Debug().Msgf("Post response: %#v", response)

	return response.Object, nil
}

// Stream
func (c *QApiClient) Stream() error {
	stream, err := c.apiClient.Stream(c.ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	//go c.send(stream)
	return c.receive(stream)
	//return c.RunTest(ctx)
	//go c.RunTest(ctx)

	// return nil
}

// Receive ...
func (c *QApiClient) receive(stream proto.QPBxApi_StreamClient) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Info().Msgf("Client connected: %+v", c)
			log.Info().Msgf("Goroutines: %d", runtime.NumGoroutine())
			return errors.Cause(err)
		}
		if err != nil {
			return errors.Cause(err)
		}

		log.Info().Msgf("==> Stream: [%T] %+v", msg, msg)
	}
}
