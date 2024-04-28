package account

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QAccount struct {
	AccountId string
	Name      string
}

func NewAccount(options ...AccountOption) *QAccount {
	h := &QAccount{
		AccountId: utils.NewID(),
		Name:      utils.NewID(),
	}

	for _, opt := range options {
		opt(h)
	}

	return h
}

func NewAccountPb(in *proto.QPBxAccount) (*QAccount, error) {
	account := NewAccount(
		WithAccountId(in.AccountId),
		WithAccountName(in.Name),
	)

	return account, nil
}

func (c *QAccount) Pb() *proto.QPBxAccount {
	return &proto.QPBxAccount{
		AccountId: c.AccountId,
		Name:      c.Name,
	}
}

func (c *QAccount) Valid() error {
	if c.AccountId == "" {
		return qlog.Error("AccountId ID not defined")
	}

	return nil
}
