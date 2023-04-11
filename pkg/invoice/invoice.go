package invoice

import (
	"github.com/mrsanchez02/go-database/pkg/invoiceheader"
	"github.com/mrsanchez02/go-database/pkg/invoiceitem"
)

// Model for invoices.
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of invoice
type Service struct {
	storage Storage
}

// NewService return a  Pointer of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new invoice
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
