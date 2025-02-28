package aplication

import (
	"API-HEX-GO/src/Products/domain"
)

type EditProduct struct {
	repo domain.IProduct
}

func NewUpdateProduct(repo domain.IProduct) *EditProduct{
	return &EditProduct{repo: repo}
}

func (cp *EditProduct) Execute(id int,product *domain.Product)error{
	return cp.repo.Update(id,product)
}