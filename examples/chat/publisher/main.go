package main

import (
	"context"
	sdk "github.com/qwibi/qwibi-go-sdk/internal/client"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/geometry"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/stream"
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

	//layer, err := client.Join(gid)
	//if err != nil {
	//	qlog.Error(err)
	//	return
	//}

	layer := geo.NewGeoObject(geometry.NewPoint(),
		//geo.WithGid("testGid"),
		geo.WithProperties([]byte("object properties")),
	)

	go func() {
		for {
			object := geo.NewGeoObject(geometry.NewPoint(),
				//geo.WithGid("testGid"),
				geo.WithProperties([]byte("object properties")),
			)

			qlog.Infof("Post object %v", object)
			//_, err = layer.Post(object)
			//if err != nil {
			//	qlog.Error(err)
			//}
			time.Sleep(3 * time.Second)
		}
	}()

	qlog.Infof("Join to layer: %+v", layer)
	err = client.Join(layer.Gid, func(m stream.QMessage) {
		qlog.Infof("Stream: request => %s", m)
	})

	if err != nil {
		qlog.Error(err)
	}
}
