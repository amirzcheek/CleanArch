package postgres

import (
	"context"
	"github.com/google/uuid"
)

func (r *Repository) AddContactToGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	return nil
}

func (r *Repository) DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	return nil
}
