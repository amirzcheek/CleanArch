package storage

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(contact *contact.Contact) (*contact.Contact, error)
	UpdateContact(ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error)
	DeleteContact(ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	ListContacts() ([]*contact.Contact, error)
	ReadContactByID(ID uuid.UUID) (contact *contact.Contact, err error)
}

type Group interface {
	CreateGroup(groupCreate *group.Group) (*group.Group, error)
	UpdateGroup(ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error)
	DeleteGroup(ID uuid.UUID) error

	GroupReader
}

type GroupReader interface {
	ListGroups() ([]*group.Group, error)
	ReadGroupByID(ID uuid.UUID) (*group.Group, error)
}

type ContactGroup interface {
	AddContactToGroup(groupID uuid.UUID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
