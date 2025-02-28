package domain

type IProduct interface {
	Save(product *Product)error
	GetAll()([]Product,error)
	GetById(id int) (*Product, error)	
	Delete(id string)error
	Update(id int,product *Product)error
}