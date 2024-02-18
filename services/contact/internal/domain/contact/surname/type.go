package surname

import "fmt"

var (
	MaxLength = 50
	ErrLength = fmt.Errorf("surname must be less than %d charachters", MaxLength)
)

type Surname string

func New(surname string) (*Surname, error) {
	if len([]rune(surname)) > MaxLength {
		return nil, ErrLength
	}
	s := Surname(surname)

	return &s, nil
}
