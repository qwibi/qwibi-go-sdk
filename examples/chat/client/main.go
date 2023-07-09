package main

import (
	"context"
	sdk "github.com/qwibi/qwibi-go-sdk"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	"time"
)

var client *sdk.QApiClient

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = sdk.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Info("Connecto to...", addr)

	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("Auth with Session... %+v", session)

	object := geo.NewGeoPoint()
	object.Properties = []byte("object properties")
	object.Feature.Properties = []byte("feature properties")

	layer, err = client.Layer("abc")
	if err != nil {
		qlog.Error(err)
		return
	}

	go func() {
		for {
			//object := geo.NewGeoPoint()
			qlog.Infof("Post object %+v", object)
			_, err = layer.Post(object)
			//if err != nil {
			//	return
			//}
			time.Sleep(3 * time.Second)
		}
	}()

	err = client.Stream(session.LayerId, func(r *proto.QPBxStreamResponse) {
		qlog.Debug(r)
	})
	if err != nil {
		panic(err)
	}

}
