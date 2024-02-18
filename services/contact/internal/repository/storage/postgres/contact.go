package postgres

import (
	"architecture_go/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

func (rep *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return nil, nil
}

func (rep *Repository) UpdateContact(ID uuid.UUID, contactUpdate contact.Contact) ([]*contact.Contact, error) {
	return nil, nil
}

func (rep *Repository) DeleteContact(ID uuid.UUID) error {
	return nil
}

func (rep *Repository) ReadContactByID(ID uuid.UUID) ([]*contact.Contact, error) {
	return nil, nil
}

func (rep *Repository) ListContact() ([]*contact.Contact, error) {
	return nil, nil
}
