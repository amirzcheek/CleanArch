package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return nil, nil
}

func (uc *UseCase) Update(contactsUpdate contact.Contact) (*contact.Contact, error) {
	return nil, nil
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return nil
}

func (uc *UseCase) List() ([]*contact.Contact, error) {
	return nil, nil
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*contact.Contact, error) {
	return nil, nil
}
