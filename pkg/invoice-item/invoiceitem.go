package invoiceItem

import "time"

type Model struct {
	ID              uint
	InvoiceHeardeID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Storage interface {
	Migrate() error
}

// Servicio de invoice item
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
