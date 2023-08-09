package qwibi

import (
	"context"
	"github.com/pkg/errors"
	auth2 "github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/metadata"
	"github.com/qwibi/qwibi-go-sdk/pkg/object"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	grpc "google.golang.org/grpc"
	"io"
	"runtime"
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
func (c *QApiClient) Auth(a auth2.QAuth) (*auth2.QSession, error) {

	switch v := a.(type) {
	case *auth2.QAnonymousAuth:
		return c.AnonymousAuth()
	case *auth2.QBasicAuth:
		return c.BasicAuth(v.Login, v.Password)
	default:
		return nil, errors.New("Unknown auth type")
	}
	// md := metadata.Pairs("token", res.Session.Token)
	// c.ctx = metadata.NewOutgoingContext(context.Background(), md)
	// return res.Session, nil
}

func (c *QApiClient) AnonymousAuth() (*auth2.QSession, error) {
	req := &proto.QPBxAuthRequest{
		Auth: &proto.QPBxAuthRequest_Anonym{
			Anonym: &proto.QPBxAnonymAuth{},
		},
	}

	res, err := c.apiClient.Auth(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	session, err := auth2.NewSessionPb(res.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	c.ctx = metadata.SetContextToken(c.ctx, session.Token)

	return session, nil
}

func (c *QApiClient) BasicAuth(login string, password string) (*auth2.QSession, error) {
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

	session, err := auth2.NewSessionPb(res.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	c.ctx = metadata.SetContextToken(c.ctx, session.Token)

	return session, nil
}

//// Stream ...
//func (c *QApiClient) Stream(gid string, handler func(m stream.QMessage)) error {
//	req := &proto.QPBxJoinRequest{
//		Gid: gid,
//	}
//
//	client, err := c.apiClient.Stream(c.ctx, req, grpc.EmptyCallOption{})
//	if err != nil {
//		return qlog.Error(err)
//	}
//
//	for {
//		msg, err := client.Recv()
//		if err == io.EOF {
//			//qlog.Infof("Client connected: %+v", c)
//			//qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
//			return qlog.Error(err)
//		}
//		if err != nil {
//			return qlog.Error(err)
//		}
//
//		qlog.Infof("==> Stream: [%T] %+v", msg, msg)
//
//		handler(stream.QMessage{})
//	}
//}

// Post ...
func (c *QApiClient) Post(object *object.QGeoObject) error {
	req := &proto.QPBxPostRequest{
		Object: object.Pb(),
	}

	_, err := c.apiClient.Post(c.ctx, req)

	return err
}

//// Get ...
//func (c *QApiClient) Get(gid string) ([]qwibi.QGeoObject, error) {
//	req := &proto.QPBxGetRequest{
//		Gid: gid,
//	}
//
//	res, err := c.apiClient.Get(c.ctx, req)
//	if err != nil {
//		return nil, qlog.Error(err)
//	}
//
//	var objects []qwibi.QGeoObject
//	for _, objectPb := range res.Objects {
//		object, err := qwibi.NewGeoObjectPb(objectPb)
//		if err != nil {
//			return nil, qlog.Error(err)
//		}
//		objects = append(objects, object)
//	}
//
//	return objects, nil
//}

func (c *QApiClient) Bot(gid string, handler func(request *command.QRequest)) error {
	client, err := c.apiClient.Bot(c.ctx, grpc.EmptyCallOption{})
	if err != nil {
		return qlog.Error(err)
	}

	for {
		in, err := client.Recv()
		if err == io.EOF {
			qlog.Infof("Client connected: %+v", c)
			qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
			return qlog.Error(err)
		}
		if err != nil {
			return qlog.Error(err)
		}

		qlog.Infof("==> Bot: [%T] %+v", in, in)

		cmd := &command.QRequest{
			Path: gid,
		}

		handler(cmd)
	}
}

// Stream ...
func (c *QApiClient) Stream(gid string, handler func(event event.QEvent)) error {
	req := &proto.QPBxStreamRequest{
		Gid: gid,
	}

	client, err := c.apiClient.Stream(c.ctx, req, grpc.EmptyCallOption{})
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
		event, err := event.NewEventPb(in.Event)
		if err != nil {
			return qlog.Error(err)
		}

		handler(event)
	}
}

//func (c *QApiClient) Stream(gid string, handler func(m stream.QMessage)) error {
//	//qlog.Infof("Start stream... %s", gid)
//	stream, err := c.apiClient.Stream(c.ctx)
//	if err != nil {
//		return qlog.Error(err)
//	}
//
//	for {
//		msg, err := stream.Recv()
//		if err == io.EOF {
//			//qlog.Infof("Client connected: %+v", c)
//			//qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
//			return qlog.Error(err)
//		}
//		if err != nil {
//			return qlog.Error(err)
//		}
//
//		qlog.Infof("==> Stream: [%T] %+v", msg, msg)
//
//		handler(stream.QMessage{})
//	}
//
//}

//func (c *QApiClient) Stream(gid string, h func(request *command.QRequest)) error {
//	client, err := c.apiClient.Stream(c.ctx, grpc.EmptyCallOption{})
//	if err != nil {
//		return qlog.Error(err)
//	}
//
//	for {
//		in, err := client.Recv()
//		if err == io.EOF {
//			qlog.Infof("Client connected: %+v", c)
//			qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
//			return qlog.Error(err)
//		}
//		if err != nil {
//			return qlog.Error(err)
//		}
//
//		qlog.Infof("==> Bot: [%T] %+v", in, in)
//
//		cmd := &command.QRequest{
//			Path: "/kkkk",
//		}
//
//		h(cmd)
//	}
//}

//// / Stream
//func (c *QApiClient) Stream() error {
//	qlog.Infof("Start stream")
//	stream, err := c.apiClient.Stream(c.ctx)
//	if err != nil {
//		return qlog.Error(err)
//	}
//
//	return c.receive(stream)
//}
//
//// Receive ...
//func (c *QApiClient) receive(stream proto.QPBxApi_StreamClient) error {
//	//x, y := 37.33172861, -122.03068446
//
//	point := qwibi.NewGeoPoint()
//
//	for {
//		msg, err := stream.Recv()
//		if err == io.EOF {
//			qlog.Infof("Client connected: %+v", c)
//			qlog.Infof("Goroutines: %d", runtime.NumGoroutine())
//			return errors.Cause(err)
//		}
//		if err != nil {
//			return errors.Cause(err)
//		}
//
//		qlog.Infof("==> Stream: [%T] %+v", msg, msg)
//
//		c.Post(point)
//
//	}
//}