package sdk

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/geo"
	"testing"

	"github.com/qwibi/qwibi-go-sdk/qlog"
)

var client *QApiClient

func init() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	c, err := NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	client = c
}

func TestAnonymousAuth(t *testing.T) {
	qlog.Debug(t.Name())
	res, err := client.AnonymousAuth()
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Anonymous auth response: %#v", res)
}

func TestBasicAuth(t *testing.T) {
	qlog.Debug(t.Name())
	res, err := client.BasicAuth("user", "password")
	if err != nil {
		t.Fatal(err)
	}
	qlog.Infof("Basic auth response: %#v", res)
}

func TestPost(t *testing.T) {
	object := geo.NewGeoPoint()

	qlog.Infof("Post request: %#v", object)

	res, err := client.Post(object)
	if err != nil {
		t.Fatal(err)
	}

	qlog.Infof("Post response: %#v", res)
}

//func TestJoinWithoutGid(t *testing.T) {
//	qlog.Debug(t.Name())
//	layer, err := client.Join("")
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Layer response: %#v", layer)
//}

//func TestJoinWithGid(t *testing.T) {
//	qlog.Debug(t.Name())
//	layer, err := client.Join("dXIUiEFyn9WUnh5q")
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Join response: %#v", layer)
//}
//
////func TestStream(t *testing.T) {
////	log.Info().Msg("Test stream connection")
////	err := client.Stream()
////	if err != nil {
////		t.Fatal(err)
////	}
////}
