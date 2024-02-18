package postgres

import "github.com/google/uuid"

func AddContactsToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	return nil
}

func DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	return nil
}
