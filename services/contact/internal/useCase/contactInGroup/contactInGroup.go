package contactInGroup

import (
	"context"
	"github.com/google/uuid"
)

func (u *UseCase) AddContactToGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	err := u.adapterStorage.AddContactToGroup(ctx, groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}
func (u *UseCase) DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	err := u.adapterStorage.DeleteContactFromGroup(ctx, groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}
