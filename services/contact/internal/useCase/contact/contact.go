package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(contact *contact.Contact) (*contact.Contact, error) {
	createdContact, err := uc.adapterStorage.CreateContact(contact)
	if err != nil {
		return nil, err
	}
	return createdContact, nil
}

func (uc *UseCase) Update(ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error) {
	updatedContact, err := uc.adapterStorage.UpdateContact(ID, contactUpdate)
	if err != nil {
		return nil, err
	}
	return updatedContact, nil
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	err := uc.adapterStorage.DeleteContact(ID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) List() ([]*contact.Contact, error) {
	contacts, err := uc.adapterStorage.ListContacts()
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*contact.Contact, error) {
	readContact, err := uc.adapterStorage.ReadContactByID(ID)
	if err != nil {
		return nil, err
	}
	return readContact, nil
}
