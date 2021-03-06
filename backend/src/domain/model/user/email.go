package user

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(v string) (Email, error) {
	var emailValidPattern = `^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	var reEmail = regexp.MustCompile(emailValidPattern)
	// Email is required
	if v == "" {
		return "", errors.New("メールアドレスの入力は必須です")
	}

	// Email Regexp
	if !reEmail.MatchString(v) {
		return "", errors.New("メールアドレスが正しくありません")
	}

	emailEntity := Email(v)

	return emailEntity, nil
}

func (email Email) Value() string {
	return string(email)
}
