package response

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/auth/account"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth/session"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

type QAuthResponse struct {
	RequestId string
	Session   *session.QSession
	Account   *account.QAccount
}

func NewAuthResponse(requestId string, session *session.QSession, account *account.QAccount) (*QAuthResponse, error) {
	res := &QAuthResponse{
		RequestId: requestId,
		Session:   session,
		Account:   account,
	}

	return res, res.Valid()
}

func NewAuthResponsePb(in *proto.QPBxAuthResponse) (*QAuthResponse, error) {
	if in == nil {
		return nil, qlog.Error("bad parameter type nil")
	}

	session, err := session.NewSessionPb(in.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	account, err := account.NewAccountPb(in.Account)
	if err != nil {
		return nil, qlog.Error(err)
	}

	return NewAuthResponse(in.RequestId, session, account)
}

func (c *QAuthResponse) Pb() *proto.QPBxAuthResponse {
	return &proto.QPBxAuthResponse{
		RequestId: c.RequestId,
		Session:   c.Session.Pb(),
		Account:   c.Account.Pb(),
	}
}

func (c *QAuthResponse) Valid() error {
	if c.Session == nil {
		return qlog.Error("session is not defined")
	}

	if c.Account == nil {
		return qlog.Error("account is not defined")
	}

	return nil
}
