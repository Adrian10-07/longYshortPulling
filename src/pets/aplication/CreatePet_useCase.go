package aplication

import "API-HEX-GO/src/pets/domain"

type CreatePet struct {
	repo domain.IPet
}

func NewCreatePet(repo domain.IPet) *CreatePet {
	return &CreatePet{repo: repo}
}

func (cp *CreatePet) Execute(p domain.Pet) error {
	return cp.repo.Save(&p)
}
