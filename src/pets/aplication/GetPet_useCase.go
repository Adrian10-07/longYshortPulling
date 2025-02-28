package aplication

import "API-HEX-GO/src/pets/domain"

type GetPet struct {
	repo domain.IPet
}

func NewGetPet(repo domain.IPet) *GetPet {
	return &GetPet{repo: repo}
}

func (cp *GetPet) Execute() ([]domain.Pet, error) {
	return cp.repo.GetAll()
}
