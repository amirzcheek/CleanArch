package postgres

import "github.com/google/uuid"

func (r *Repository) AddContactsToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	return nil
}

func (r *Repository) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	return nil
}
