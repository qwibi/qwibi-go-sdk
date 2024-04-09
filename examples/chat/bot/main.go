package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	sdkCommand "github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/layer"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/qwibi"
	"strings"
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
		"/help":    "help description",
		"/message": "send a message",
	}

	result, err := client.Layer(
		layer.WithLayerGid("chat"),
		layer.WithLayerCommands(commands),
	)
	if err != nil {
		qlog.Error(err)
		return
	}
	qlog.Infof("Response: %+v", result)

	token, err := client.Token(result.Layer.LayerId)
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

	err = bot.Subscribe(func(requestId string, layerId string, command *sdkCommand.QCommand) {
		qlog.Infof("bot command: %+v", command)

		status := fmt.Sprintf("unknown command: %s", command.Command)
		response := sdkCommand.NewResponse(status)

		switch command.Command {
		case "/help":
			jsonData, err := json.Marshal(result.Layer.Commands)
			if err != nil {
				qlog.Error(err)
			}
			response = sdkCommand.NewResponse(string(jsonData))

		case "/message":
			response = messageHandler(command.Args)
		}

		qlog.Infof("bot response: %+v", response)

		err = bot.Publish(requestId, layerId, response)
		if err != nil {
			qlog.Error(err)
			return
		}
	})

	if err != nil {
		qlog.Error(err)
	}
}

func messageHandler(args []string) *sdkCommand.QResponse {
	text := strings.Join(args, " ")
	status := fmt.Sprintf(text)
	response := sdkCommand.NewResponse(status)
	return response
}
