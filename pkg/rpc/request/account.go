package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/auth/account"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QAccountRequest struct {
	RequestId string
	Account   *account.QAccount
}

func NewAccountRequest(account *account.QAccount) (*QAccountRequest, error) {
	req := &QAccountRequest{
		RequestId: utils.RequestId(),
		Account:   account,
	}

	return req, account.Valid()
}

func NewAccountRequestPb(in *proto.QPBxAccountRequest) (*QAccountRequest, error) {
	account, err := account.NewAccountPb(in.Account)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewAccountRequest(account)
}

func (c *QAccountRequest) Pb() *proto.QPBxAccountRequest {
	return &proto.QPBxAccountRequest{
		RequestId: c.RequestId,
		Account:   c.Account.Pb(),
	}
}

func (c *QAccountRequest) Message() protobuf.Message {
	return c.Pb()
}

func (c *QAccountRequest) ProtoReflect() protoreflect.Message {
	return c.Pb().ProtoReflect()
}

func (c *QAccountRequest) Event() *proto.QPBxEvent {
	return &proto.QPBxEvent{
		Type: &proto.QPBxEvent_AccountRequest{
			AccountRequest: c.Pb(),
		},
	}
}
