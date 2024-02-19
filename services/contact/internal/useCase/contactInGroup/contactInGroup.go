package contactInGroup

import (
	"github.com/google/uuid"
)

func (u *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	err := u.adapterStorage.AddContactToGroup(groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}
func (u *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	err := u.adapterStorage.DeleteContactFromGroup(groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}
