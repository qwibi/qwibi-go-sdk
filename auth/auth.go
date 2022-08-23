package auth

type QAuth interface {
	Valid() error
}
