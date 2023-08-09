package main

import (
	"context"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/event"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
)

var client *qwibi.QApiClient

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = qwibi.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Info("Connecto to...", addr)

	session, err := client.Auth(&auth.QAnonymousAuth{})
	if err != nil {
		panic(err)
	}
	qlog.Infof("Auth with Session... %+v", session)

	//object := qwibi.NewGeoPoint()
	//object.Pr
	//operties = []byte("object properties")
	//object.Feature.Properties = []byte("feature properties")

	//layer, err := client.Layer("chat")
	//if err != nil {
	//	qlog.Error(err)
	//	return
	//}

	gid := "myID"
	qlog.Infof("Subscribe stream for gid: %s", gid)
	err = client.Stream(gid, func(event event.QEvent) {
		qlog.Infof("Event: %+v", event)
	})

	if err != nil {
		qlog.Error(err)
	}
}
