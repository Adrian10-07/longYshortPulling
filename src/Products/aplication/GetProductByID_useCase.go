package aplication

import "API-HEX-GO/src/Products/domain"

type GetProductById struct {
	repo domain.IProduct
}

func NewGetProductById(repo domain.IProduct) *GetProductById {
	return &GetProductById{repo: repo}
}

func (cp *GetProductById) Execute(id int) (*domain.Product, error) {
	product, err := cp.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
