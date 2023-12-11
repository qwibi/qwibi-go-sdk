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

	req := &proto.QPBxLayerRequest{
		LayerId:    g.LayerId(),
		Public:     g.Public(),
		Properties: g.Properties(),
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

//func (c *QApiClient) Bot(layerId string) (*geoBot, error) {
//	client, err := c.apiClient.Bot(c.ctx, grpc.EmptyCallOption{})
//	if err != nil {
//		return nil, qlog.Error(err)
//	}
//
//}

//// Stream ...
//func (c *QApiClient) Stream(gid string, handler func(m stream.QMessage)) error {
//	req := &proto.QPBxJoinRequest{
//		gid: gid,
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

//// Get ...
//func (c *QApiClient) Get(gid string) ([]qwibi.QGeoObject, error) {
//	req := &proto.QPBxGetRequest{
//		gid: gid,
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
//		c.Node(point)
//
//	}
//}
