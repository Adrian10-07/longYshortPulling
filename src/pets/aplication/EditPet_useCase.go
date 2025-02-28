package aplication

import (
	"API-HEX-GO/src/pets/domain"
)

type EditPet struct {
	repo domain.IPet
}

func NewEditPet(repo domain.IPet) *EditPet {
	return &EditPet{repo: repo}
}

func (cp *EditPet) Execute(id string, pet *domain.Pet) error {
	return cp.repo.Update(id, pet)
}
