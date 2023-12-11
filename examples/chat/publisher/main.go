package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo"
	"github.com/qwibi/qwibi-go-sdk/pkg/geo/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
	"math/rand"
	"time"
)

var client *qwibi.QApiClient

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = qwibi.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Info("connecto to...", addr)

	a := &auth.QAnonymousAuth{
		//Token: "111",
		//Token: "OZbHVLm0Eb2Tyhq6",
	}

	session, err := client.Auth(a)
	if err != nil {
		panic(err)
	}
	qlog.Infof("auth with session... %+v", session)

	layer, err := client.Layer(
		layer.WithLayerGid("mrt"),
		//layer.WithLayerPublic(true),
	)
	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("layer: %+v", layer)

	go func() {
		for {
			point := geo.NewGeoPoint(
				geo.WithPointGid("me"),
				geo.WithPointCoordinates(1.1+rand.Float64(), 2.2+rand.Float64()),
			)

			qlog.Infof("post[%s]: %+v", layer.Gid(), point.Geometry())
			err = layer.Post(point)
			if err != nil {
				qlog.Fatal(err)
			}

			time.Sleep(1000 * time.Millisecond)
		}
	}()

	err = layer.Stream(func(event event.QEvent) {
		qlog.Infof("event[%s]: [%T] %+v", layer.Gid(), event, event)
	})

	if err != nil {
		qlog.Error(err)
	}
}
