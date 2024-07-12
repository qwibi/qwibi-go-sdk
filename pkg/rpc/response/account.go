package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/auth/account"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QAccountResponse struct {
	RequestId string
	Account   *account.QAccount
}

func NewAccountResponse(requestId string, account *account.QAccount) (*QAccountResponse, error) {
	if requestId == "" {
		return nil, qlog.Error("request ID not defined")
	}

	if account == nil {
		return nil, qlog.Error("account is not defined")
	}

	res := &QAccountResponse{
		RequestId: requestId,
		Account:   account,
	}

	return res, nil
}

func NewAccountResponsePb(in *proto.QPBxAccountResponse) (*QAccountResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	account, err := account.NewAccountPb(in.Account)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewAccountResponse(in.RequestId, account)
}

func (c *QAccountResponse) Pb() *proto.QPBxAccountResponse {
	return &proto.QPBxAccountResponse{
		RequestId: c.RequestId,
		Account:   c.Account.Pb(),
	}
}

func (c *QAccountResponse) Message() protobuf.Message {
	return c.Pb()
}

func (c *QAccountResponse) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QAccountResponse) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_AccountResponse{
			AccountResponse: c.Pb(),
		},
	}
}
