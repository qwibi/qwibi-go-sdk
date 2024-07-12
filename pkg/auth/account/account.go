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

func NewAccount(options ...AccountOption) (*QAccount, error) {
	account := &QAccount{
		AccountId: utils.NewID(),
		Name:      utils.NewID(),
	}

	for _, opt := range options {
		opt(account)
	}

	return account, account.Valid()
}

func NewAccountPb(in *proto.QPBxAccount) (*QAccount, error) {
	return NewAccount(
		WithAccountId(in.AccountId),
		WithAccountName(in.Name),
	)
}

func (c *QAccount) Pb() *proto.QPBxAccount {
	return &proto.QPBxAccount{
		AccountId: c.AccountId,
		Name:      c.Name,
	}
}

func (c *QAccount) Valid() error {
	if c.AccountId == "" {
		return qlog.Error("account ID is not defined")
	}

	if c.Name == "" {
		return qlog.Error("account name is not defined")
	}

	return nil
}
