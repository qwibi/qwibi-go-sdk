package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/command"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/rpc/request"
	sdkResponse "github.com/qwibi/qwibi-go-sdk/pkg/rpc/response"
)

// Command ...
func (c *QApiClient) Command(layerId string, command *command.QCommand) (*sdkResponse.QCommandResponse, error) {

	req, err := request.NewCommandRequest(layerId, command)
	if err != nil {
		return nil, qlog.Error(err)
	}

	in, err := c.apiClient.Command(c.ctx, req.Pb())
	if err != nil {
		return nil, qlog.Error(err)
	}

	res, err := sdkResponse.NewCommandResponsePb(in)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return res, nil
}
