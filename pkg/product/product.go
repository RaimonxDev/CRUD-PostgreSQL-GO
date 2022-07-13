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
	Migrate() error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GeyByID(uint) (*Model, error)
	//Delete(uint) error
}

// Servicio de productos
type Service struct {
	// Embebed Interface
	storage Storage
}

// New Service return un puntero de Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Ejecutamos el metodo migrate de Storage desde el servicios
// Debido a que embebimos storage
func (s Service) Migrate() error {
	// Como migrate tambien devuelve un error lo returnamos sin mas codigo
	return s.storage.Migrate()
}
