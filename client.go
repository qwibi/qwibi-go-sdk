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

	md := metadata.Pairs("token", res.Token)
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
func (c *QApiClient) Join(layerID string) (*geo.QLayer, error) {
	joinRequest, err := request.NewJoinRequest(layerID)
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

	layer, err := geo.NewLayerPb(joinResponse.Layer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return layer, nil
}

// Join ...
func (c *QApiClient) Post(object *geo.QGeoObject) (*geo.QGeoObject, error) {
	request, err := request.NewPostRequest(object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	requestPb, err := request.Pb()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Debug().Msgf("RPC request: %s", requestPb.String())
	responsePb, err := c.apiClient.Post(c.ctx, requestPb)
	if err != nil {
		// st, ok := status.FromError(err)
		// if ok {
		// 	fmt.Printf("!!!!! %v", st.Message())
		// }
		err := errors.New(fmt.Sprint(err))
		log.Error().Stack().Err(err).Msg("")
		return nil, errors.WithStack(err)
	}

	response, err := response.NewPostResponsePb(responsePb)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response.Object, nil
}

//// Stream
//func (c *QApiClient) Stream(ctx context.Context ) error {
//	streamRequest, err := request.NewStreamRequest()
//	if err != nil {
//		return errors.WithStack(err)
//	}
//
//	pb, err := streamRequest.Pb()
//	if err != nil {
//		return errors.WithStack(err)
//	}
//
//	stream, err := c.apiClient.Stream(ctx, pb)
//	if err != nil {
//		return errors.WithStack(err)
//	}
//
//	//go c.send(stream)
//	go c.receive(stream)
//	//return c.RunTest(ctx)
//	//go c.RunTest(ctx)
//
//	return nil
//}

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

		log.Info().Msgf("==> Stream: [%T] %+v", msg, msg)
	}
}
