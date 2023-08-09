package qwibi

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"testing"
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

//func TestPostLayer(t *testing.T) {
//	object := qwibi.NewGeoPoint()
//
//	qlog.Infof("Post request: %#v", object)
//	res, err := client.Post(object)
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Post response: %#v", res)
//
//	qlog.Infof("Update request: %#v", object)
//	res2, err := client.Post(object)
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Update response: %#v", res2)
//}

//func TestGetObjects(t *testing.T) {
//	testId := fmt.Sprint(rand.Int())
//	object := qwibi.NewGeoPoint()
//	object.Gid = testId
//
//	qlog.Infof("Post request: %#v", object)
//	res, err := client.Post(object)
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Post response: %#v", res)
//
//	res2, err := client.Get(testId)
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Get response: %#v", res2)
//
//	//qlog.Infof("Update request: %#v", object)
//	//res2, err := client.Post(object)
//	//if err != nil {
//	//	t.Fatal(err)
//	//}
//	//qlog.Infof("Update response: %#v", res2)
//}

//func TestPostPoint(t *testing.T) {
//	object := qwibi.NewGeoPoint()
//
//	qlog.Infof("Post request: %#v", object)
//
//	res, err := client.Post(object)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	qlog.Infof("Post response: %#v", res)
//}

//func TestJoinWithoutGid(t *testing.T) {
//	qlog.Debug(t.Name())
//	layer, err := client.Stream("")
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Layer response: %#v", layer)
//}

//func TestJoinWithGid(t *testing.T) {
//	qlog.Debug(t.Name())
//	layer, err := client.Stream("dXIUiEFyn9WUnh5q")
//	if err != nil {
//		t.Fatal(err)
//	}
//	qlog.Infof("Stream response: %#v", layer)
//}
//
////func TestStream(t *testing.T) {
////	log.Info().Msg("Test stream connection")
////	err := client.Stream()
////	if err != nil {
////		t.Fatal(err)
////	}
////}
