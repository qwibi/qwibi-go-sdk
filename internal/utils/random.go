package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

func Random(k int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s, _ := gonanoid.Generate(chars, k)
	return s
}

func NewID() string {
	return Random(16)
}
