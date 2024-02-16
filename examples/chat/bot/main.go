package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
	sdkRequest "github.com/qwibi/qwibi-go-sdk/pkg/rpc/request"
)

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()
	var client, err = qwibi.NewClient(ctx, addr)
	if err != nil {
		panic(err)
	}
	qlog.Infof("connect to: %s", addr)

	session, err := client.Auth(&auth.QBasicAuth{
		Login:    "maggnus",
		Password: "test",
	})
	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("auth session: %+v", session)

	commands := map[string]string{
		"/help":   "help description",
		"/locate": "locate something",
	}

	layer, err := client.Layer(
		layer.WithLayerGid("chat"),
		layer.WithLayerCommands(commands),
	)
	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("layer: %+v", layer)

	token, err := client.Token(layer.LayerId)
	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("token: %+v", token)

	bot, err := client.Bot(token)
	if err != nil {
		qlog.Error(err)
		return
	}

	err = bot.Subscribe(func(request *sdkRequest.QCommandRequest) {
		qlog.Infof("bot request: %+v", request)

		var res *command.QResponse
		commandName := request.Command.Command
		if _, ok := layer.Commands[commandName]; !ok {
			status := fmt.Sprintf("command not found: %s", commandName)
			qlog.Warnf(status)
			res = command.NewResponse(status)
		} else {
			switch commandName {
			case "/help":
				jsonData, err := json.Marshal(layer.Commands)
				if err != nil {
					qlog.Error(err)
				}
				res = command.NewResponse(string(jsonData))
			default:
				status := fmt.Sprintf("unknown command: %s", commandName)
				qlog.Warnf(status)
				res = command.NewResponse(status)
			}
		}

		qlog.Infof("bot response: %+v", res)

		err = bot.Publish(request.RequestId, request.LayerId, res)
		if err != nil {
			qlog.Error(err)
			return
		}
	})

	if err != nil {
		qlog.Error(err)
	}
}
