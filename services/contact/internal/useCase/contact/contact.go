package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func (uc *UseCase) Create(ctx context.Context, contact *contact.Contact) (*contact.Contact, error) {
	log.Printf("Starting creation of contact %s", contact.FullName())

	traceID := uuid.New().String()
	log.Printf("[%s] Started tracing for creating contact", traceID)

	existingContact, err := uc.adapterStorage.ReadContactByFullName(ctx, contact.FullName())
	if err != nil {
		return nil, err
	}
	if existingContact != nil {
		return nil, fmt.Errorf("contact with name %s already exists", contact.FullName())
	}

	createdContact, err := uc.adapterStorage.CreateContact(ctx, contact)
	if err != nil {
		return nil, err
	}

	log.Printf("Contact %s created successfully", contact.FullName())

	return createdContact, nil
}

func (uc *UseCase) Update(ctx context.Context, ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error) {
	updatedContact, err := uc.adapterStorage.UpdateContact(ctx, ID, contactUpdate)
	if err != nil {
		return nil, err
	}
	return updatedContact, nil
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	err := uc.adapterStorage.DeleteContact(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) List(ctx context.Context) ([]*contact.Contact, error) {
	contacts, err := uc.adapterStorage.ListContacts(ctx)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (uc *UseCase) ReadByID(ctx context.Context, ID uuid.UUID) (*contact.Contact, error) {
	readContact, err := uc.adapterStorage.ReadContactByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return readContact, nil
}
