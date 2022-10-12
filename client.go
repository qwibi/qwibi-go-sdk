package sdk

import (
	"context"
	"io"
	"runtime"

	"github.com/pkg/errors"
	"github.com/qwibi/qwibi-go-sdk/auth"
	"github.com/qwibi/qwibi-go-sdk/feature"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"github.com/qwibi/qwibi-go-sdk/geometry"
	"github.com/qwibi/qwibi-go-sdk/metadata"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"github.com/qwibi/qwibi-go-sdk/qlog"
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

// Auth ...
func (c *QApiClient) Auth(a auth.QAuth) (*auth.QSession, error) {

	switch v := a.(type) {
	case *auth.QAnonymousAuth:
		return c.AnonymousAuth()
	case *auth.QBasicAuth:
		return c.BasicAuth(v.Login, v.Password)
	default:
		return nil, errors.New("Unknown auth type")
	}
	// md := metadata.Pairs("token", res.Session.Token)
	// c.ctx = metadata.NewOutgoingContext(context.Background(), md)
	// return res.Session, nil
}

func (c *QApiClient) AnonymousAuth() (*auth.QSession, error) {
	req := &proto.QPBxAuthRequest{
		Auth: &proto.QPBxAuthRequest_Anonym{
			Anonym: &proto.QPBxAnonymAuth{},
		},
	}

	res, err := c.apiClient.Auth(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	session, err := auth.NewSessionPb(res.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	c.ctx = metadata.SetContextToken(c.ctx, session.Token)

	return session, nil
}

func (c *QApiClient) BasicAuth(login string, password string) (*auth.QSession, error) {
	req := &proto.QPBxAuthRequest{
		Auth: &proto.QPBxAuthRequest_Basic{
			Basic: &proto.QPBxBasicAuth{
				Login:    login,
				Password: password,
			},
		},
	}

	res, err := c.apiClient.Auth(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	session, err := auth.NewSessionPb(res.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	c.ctx = metadata.SetContextToken(c.ctx, session.Token)

	return session, nil
}

// Join ...
func (c *QApiClient) Join(gid string) (geo.QGeoObject, error) {
	req := &proto.QPBxJoinRequest{
		Gid: gid,
	}

	res, err := c.apiClient.Join(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	object, err := geo.NewGeoObjectPb(res.Object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return object, nil
}

// Join ...
func (c *QApiClient) Post(object geo.QGeoObject) (geo.QGeoObject, error) {
	req := &proto.QPBxPostRequest{
		Object: object.Pb(),
	}

	res, err := c.apiClient.Post(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	object, err = geo.NewGeoObjectPb(res.Object)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return object, nil
}

// Stream
func (c *QApiClient) Stream() error {
	stream, err := c.apiClient.Stream(c.ctx)
	if err != nil {
		return qlog.Error(err)
	}

	return c.receive(stream)
}

// Receive ...
func (c *QApiClient) receive(stream proto.QPBxApi_StreamClient) error {
	x, y := 37.33172861, -122.03068446
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			qlog.Infof("Client connected: %+v", c)
			qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
			return errors.Cause(err)
		}
		if err != nil {
			return errors.Cause(err)
		}

		qlog.Infof("==> Stream: [%T] %+v", msg, msg)

		c.Post(&geo.QGeoPoint{
			Gid: "123456",
			Feature: &feature.QPointFeature{
				Geometry: &geometry.QPoint{
					Coordinates: []float64{x, y},
				},
			},
		})

	}
}
