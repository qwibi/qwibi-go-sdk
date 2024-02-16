package utils

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"regexp"
)

func IsValidLogin(login string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !re.MatchString(login) {
		return qlog.Error("allow only alphanumeric characters and underscores")
	}

	return nil
}

func IsValidPassword(password string) error {
	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`)
	if !re.MatchString(password) {
		return qlog.Error("require at least 8 characters, including at least one uppercase letter, one lowercase letter, and one digit")
	}

	return nil
}
