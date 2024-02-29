package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"github.com/google/uuid"
)

type Contact interface {
	Create(ctx context.Context, contact *contact.Contact) (*contact.Contact, error)
	Update(ctx context.Context, ID uuid.UUID, contactUpdate contact.Contact) (*contact.Contact, error)
	Delete(ctx context.Context, ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	List(ctx context.Context) ([]*contact.Contact, error)
	ReadByID(ctx context.Context, ID uuid.UUID) (contact *contact.Contact, err error)
}

type Group interface {
	Create(ctx context.Context, groupCreate *group.Group) (*group.Group, error)
	Update(ctx context.Context, ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	GroupReader
}

type GroupReader interface {
	List(ctx context.Context) ([]*group.Group, error)
	ReadByID(ctx context.Context, ID uuid.UUID) (*group.Group, error)
}

type ContactGroup interface {
	AddContactToGroup(ctx context.Context, groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error
}
