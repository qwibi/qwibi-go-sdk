package auth

type AccountOption func(config *QAccount)

func WithAccountId(accountID string) AccountOption {
	return func(c *QAccount) {
		c.AccountId = accountID
	}
}

func WithAccountName(name string) AccountOption {
	return func(c *QAccount) {
		c.Name = name
	}
}

//func AccountProperties(properties []byte) AccountOption {
//	return func(c *QAccount) {
//		c.properties = properties
//	}
//}
