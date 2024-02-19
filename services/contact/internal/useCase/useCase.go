package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Contact interface {
	Create(contact *contact.Contact) (*contact.Contact, error)
	Update(ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error)
	Delete(ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	List() ([]*contact.Contact, error)
	ReadByID(ID uuid.UUID) (contact *contact.Contact, err error)
}

type Group interface {
	Create(groupCreate *group.Group) (*group.Group, error)
	Update(ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error)
	Delete(ID uuid.UUID) error
	GroupReader
}

type GroupReader interface {
	List() ([]*group.Group, error)
	ReadByID(ID uuid.UUID) (*group.Group, error)
}

type ContactGroup interface {
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
