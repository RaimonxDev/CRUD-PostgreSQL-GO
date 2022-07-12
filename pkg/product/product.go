package product

import "time"

// Model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAT   time.Time
}

// Slices de modelos
type Models []*Model

// Todos los metodos del CRUD
type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GeyByID(uint) (*Model, error)
	Delete(uint) error
}
