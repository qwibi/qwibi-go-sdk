package request

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/auth/account"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type QAccountRequest struct {
	RequestId string
	Account   *account.QAccount
}

func NewAccountRequest(options ...account.AccountOption) (*QAccountRequest, error) {
	account := account.NewAccount(options...)
	req := &QAccountRequest{
		RequestId: utils.RequestId(),
		Account:   account,
	}

	return req, account.Valid()
}

func NewAccountRequestPb(in *proto.QPBxAccountRequest) (*QAccountRequest, error) {
	return NewAccountRequest(
		account.WithAccountId(in.AccountId),
		account.WithAccountName(in.Name),
	)
}

func (c *QAccountRequest) Pb() *proto.QPBxAccountRequest {
	return &proto.QPBxAccountRequest{
		AccountId: c.Account.AccountId,
		Name:      c.Account.Name,
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
