package account

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/pkg/utils"
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

func (c *QAccount) Valid() error {
	if c.AccountId == "" {
		return qlog.Error("Account ID not defined")
	}

	return nil
}
