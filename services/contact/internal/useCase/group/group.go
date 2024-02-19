package group

import (
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	createdGroup, err := uc.adapterStorage.CreateGroup(groupCreate)
	if err != nil {
		return nil, err
	}
	return createdGroup, nil
}

func (uc *UseCase) Update(ID uuid.UUID, groupUpdate *group.Group) (*group.Group, error) {
	updatedGroup, err := uc.adapterStorage.UpdateGroup(ID, groupUpdate)
	if err != nil {
		return nil, err
	}
	return updatedGroup, nil
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	err := uc.adapterStorage.DeleteGroup(ID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) List() ([]*group.Group, error) {
	groups, err := uc.adapterStorage.ListGroups()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	readGroup, err := uc.adapterStorage.ReadGroupByID(ID)
	if err != nil {
		return nil, err
	}
	return readGroup, nil
}
