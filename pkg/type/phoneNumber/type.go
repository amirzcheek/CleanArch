package phoneNumber

import (
	"errors"
	"regexp"
)

type PhoneNumber struct {
	value string
}

func New(phone string) (*PhoneNumber, error) {
	re := regexp.MustCompile(`^[0-9]{11}$`)
	if !re.MatchString(phone) {
		return nil, errors.New("invalid phone number")
	}
	return &PhoneNumber{value: phone}, nil
}
