package product

import (
	"fmt"
	"strings"
	"time"
)

// Model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observation, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Slices de modelos
type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

// Todos los metodos del CRUD
type Storage interface {
	Migrate() error
	Create(*Model) error
	//Update(*Model) error
	GetAll() (Models, error)
	GeyByID(uint) (*Model, error)
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

// Ejecutamos el metodo migrate de Storage desde el servicio
// Debido a que embebimos storage
func (s *Service) Migrate() error {
	// Como migrate tambien devuelve un error lo returnamos sin mas codigo
	return s.storage.Migrate()
}

func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GeyByID(id)
}
