package session

type SessionOption func(config *QSession)

func WithSessionAccessToken(accessToken string) SessionOption {
	return func(c *QSession) {
		c.AccessToken = accessToken
	}
}

func WithSessionAccountId(accountId string) SessionOption {
	return func(c *QSession) {
		c.AccountId = accountId
	}
}
