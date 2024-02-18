package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Contact interface {
	Create(contacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(contactUpdate contact.Contact) (*contact.Contact, error)
	Delete(ID uuid.UUID) (string, error)

	ContactReader
}

type ContactReader interface {
	List() ([]*contact.Contact, error)
	ReadByID(ID uuid.UUID) (contact *contact.Contact, err error)
}

type Group interface {
	Create(group *group.Group) (*group.Group, error)
	Update(groupUpdate *group.Group) (*group.Group, error)

	ContactGroup
}

type GroupReader interface {
	List() ([]*group.Group, error)
	ReadByID(ID uuid.UUID) (*group.Group, error)
}

type ContactGroup interface {
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
