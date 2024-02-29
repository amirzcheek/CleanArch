package postgres

import (
	"architecture_go/services/contact/internal/domain/contact"
	"context"
	"github.com/google/uuid"
	"log"
)

func (r *Repository) CreateContact(ctx context.Context, contact *contact.Contact) (*contact.Contact, error) {
	log.Printf("Starting creation of contact %s in the repository", contact.FullName())

	log.Printf("Contact %s created successfully in the repository", contact.FullName())

	return contact, nil
}

func (r *Repository) UpdateContact(ctx context.Context, ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error) {
	return nil, nil
}

func (r *Repository) DeleteContact(ctx context.Context, ID uuid.UUID) error {
	return nil
}

func (r *Repository) ReadContactByID(ctx context.Context, ID uuid.UUID) (*contact.Contact, error) {
	return nil, nil
}

func (r *Repository) ListContacts(ctx context.Context) ([]*contact.Contact, error) {
	return nil, nil
}
