package name

import "fmt"

var (
	MaxLength = 50
	ErrLength = fmt.Errorf("name must be less than %d charachters", MaxLength)
)

type Name string

func New(name string) (*Name, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrLength
	}
	s := Name(name)

	return &s, nil
}
