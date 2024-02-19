package contactInGroup

import "architecture_go/services/contact/internal/useCase/adapters/storage"

type UseCase struct {
	adapterStorage storage.ContactGroup
}

func New(storage storage.ContactGroup) *UseCase {
	var u = &UseCase{
		adapterStorage: storage,
	}
	return u
}
