package auth

type AnonymousAuthOption func(config *QAnonymousAuth)

func WithAnonymousAuthToken(token string) AnonymousAuthOption {
	return func(c *QAnonymousAuth) {
		c.Token = token
	}
}
