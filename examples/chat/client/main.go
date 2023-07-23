package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	sdk "github.com/qwibi/qwibi-go-sdk/pkg/client"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/stream"
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
	//object.Feature.Properties = []byte("feature properties")

	layer, err := client.Layer("chat")
	if err != nil {
		qlog.Error(err)
		return
	}

	qlog.Infof("Join to layer: %+v", layer)
	err = client.Join(layer.Gid, func(m stream.QMessage) {
		qlog.Infof("Stream: request => %s", m)
	})

	if err != nil {
		qlog.Error(err)
	}
}
