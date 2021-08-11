package utils

type Token = string

func NewToken() Token {
	return Random(21)
}
