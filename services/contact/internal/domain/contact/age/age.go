package age

import (
	"fmt"
)

var (
	MaxAge uint8 = 250
	ErrAge       = fmt.Errorf("age must be less than %d", MaxAge)
)

type Age uint8

func New(age uint8) (*Age, error) {
	if age > MaxAge {
		return nil, ErrAge
	}
	a := Age(age)
	return &a, nil
}
