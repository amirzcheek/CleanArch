package email

import (
	"errors"
	"regexp"
)

type Email struct {
	address string
}

func New(address string) (*Email, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(address) {
		return nil, errors.New("invalid email")
	}
	return &Email{address: address}, nil
}
