package aplication

import "API-HEX-GO/src/Products/domain"

type DeleteProduct struct {
	rep domain.IProduct
}

func NewDeleteProduct(rep domain.IProduct) *DeleteProduct{
	return &DeleteProduct{
		rep: rep,
	}
}

func (cp *DeleteProduct) Execute(nombre string)error {
	return cp.rep.Delete(nombre)
}