package main

import (
	"context"
	sdk "github.com/qwibi/qwibi-go-sdk"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
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

	layer, err = client.Join("chat")
	if err != nil {
		qlog.Fatal(err)
		panic("")
	}
	qlog.Infof("Join to... %+v", "/chat")

	layer.Post("/hello")

	layers.GetObjects()

	go func() {
		for {
			qlog.TODO("Send command... /")
			time.Sleep(3 * time.Second)
		}
	}()

	err = client.Stream(func(r *proto.QPBxStreamResponse) {
		qlog.Debug(r)
	})

	if err != nil {
		qlog.Fatal(err)
		panic("")

	}

}
