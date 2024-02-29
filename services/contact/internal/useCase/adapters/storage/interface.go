package storage

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"github.com/google/uuid"
)

type Storage interface {
	Contact
	Group
	ContactGroup
}

type Contact interface {
	CreateContact(ctx context.Context, contact *contact.Contact) (*contact.Contact, error)
	UpdateContact(ctx context.Context, ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error)
	DeleteContact(ctx context.Context, ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	ListContacts(ctx context.Context) ([]*contact.Contact, error)
	ReadContactByID(ctx context.Context, ID uuid.UUID) (contact *contact.Contact, err error)
}

type Group interface {
	CreateGroup(ctx context.Context, groupCreate *group.Group) (*group.Group, error)
	UpdateGroup(ctx context.Context, ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error)
	DeleteGroup(ctx context.Context, ID uuid.UUID) error

	GroupReader
}

type GroupReader interface {
	ListGroups(ctx context.Context) ([]*group.Group, error)
	ReadGroupByID(ctx context.Context, ID uuid.UUID) (*group.Group, error)
}

type ContactGroup interface {
	AddContactToGroup(ctx context.Context, groupID uuid.UUID, contactID uuid.UUID) error
	DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error
}
