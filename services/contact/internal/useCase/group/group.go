package group

import (
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	return nil, nil
}

func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	return nil, nil
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return nil
}

func (uc *UseCase) List() ([]*group.Group, error) {
	return nil, nil
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	return nil, nil
}
