package utils

import (
	"github.com/GoWebProd/uuid7"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"sync"
)

var (
	uuid = uuid7.New()
	once sync.Once
)

func Random(k int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s, _ := gonanoid.Generate(chars, k)
	return s
}

func NewID() string {
	return Random(16)
}

func RequestId() string {
	//return strings.ToLower(Random(8))
	return uuid.Next().String()
}
