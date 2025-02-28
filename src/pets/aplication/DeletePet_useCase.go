package aplication

import "API-HEX-GO/src/pets/domain"

type DeletePet struct {
	rep domain.IPet
}

func NewDeletePet(rep domain.IPet) *DeletePet {
	return &DeletePet{
		rep: rep,
	}
}

func (cp *DeletePet) Execute(nombre string) error {
	return cp.rep.Delete(nombre)
}
