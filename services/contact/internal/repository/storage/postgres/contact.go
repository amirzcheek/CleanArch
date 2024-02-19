package postgres

import (
	"architecture_go/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

func (r *Repository) CreateContact(contact *contact.Contact) (*contact.Contact, error) {
	return nil, nil
}

func (r *Repository) UpdateContact(ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error) {
	return nil, nil
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	return nil
}

func (r *Repository) ReadContactByID(ID uuid.UUID) (*contact.Contact, error) {
	return nil, nil
}

func (r *Repository) ListContacts() ([]*contact.Contact, error) {
	return nil, nil
}
