package domain

type IPet interface {
	Save(pet *Pet)error
	GetAll()([]Pet,error)
	Delete(id string)error
	Update(id string,pet *Pet)error
	
}