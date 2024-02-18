package gender

import "errors"

type Gender struct {
	value string
}

var (
	Male   = Gender{value: "male"}
	Female = Gender{value: "female"}
)

func NewGender(gender string) (*Gender, error) {
	if gender != Male.value && gender != Female.value {
		return nil, errors.New("invalid gender")
	}
	return &Gender{value: gender}, nil
}
