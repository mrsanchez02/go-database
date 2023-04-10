package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Models slice of Model
type Models []*Model

// Storage interface that must implement a db storage.
type Storage interface {
	Migrate() error
}

// Service of invoiceheader
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceheader.
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}