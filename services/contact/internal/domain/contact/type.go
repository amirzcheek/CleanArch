package contact

import (
	"architecture_go/pkg/type/email"
	"architecture_go/pkg/type/gender"
	"architecture_go/pkg/type/phoneNumber"
	"architecture_go/services/contact/internal/domain/contact/name"
	"architecture_go/services/contact/internal/domain/contact/surname"
	"fmt"
	"github.com/google/uuid"
)

type Contact struct {
	ID          uuid.UUID
	Name        name.Name
	Surname     surname.Surname
	PhoneNumber phoneNumber.PhoneNumber

	Email  email.Email
	Gender gender.Gender
}

func NewContact(firstName name.Name, lastName surname.Surname,
	phoneNumber phoneNumber.PhoneNumber, email email.Email, gender gender.Gender) *Contact {
	return &Contact{
		ID:          uuid.New(),
		Name:        firstName,
		Surname:     lastName,
		PhoneNumber: phoneNumber,
		Email:       email,
		Gender:      gender,
	}
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.Surname, c.Name)
}
