package auth

type SessionOption func(config *QSession)

func WithSessionToken(token string) SessionOption {
	return func(c *QSession) {
		c.Token = token
	}
}

func WithSessionAccountId(accountId string) SessionOption {
	return func(c *QSession) {
		c.AccountId = accountId
	}
}
