package core

type CrudDAO interface {
	FindById()
	Insert()
	Delete()
	Update()
}
