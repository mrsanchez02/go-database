package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrIDNotFound = errors.New("product doesn't have an ID")
)

// Model for products.
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreateAt     time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreateAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Models a slice of Model
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

// Storage interface that must implement a db storage.
type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

// Service of product
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product.
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create is used to creating a product.
func (s *Service) Create(m *Model) error {
	m.CreateAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used to get all the products.
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// // GetByID is used to get a product.
// func (s *Service) GetByID(id uint) (*Model, error) {
// 	return s.storage.GetByID(id)
// }

// // Update is used to modify a product.
// func (s *Service) Update(m *Model) error {
// 	if m.ID == 0 {
// 		return ErrIDNotFound
// 	}
// 	m.UpdatedAt = time.Now()

// 	return s.storage.Update(m)
// }

// // Delete is used to delete a product.
// func (s *Service) Delete(id uint) error {
// 	if id == 0 {
// 		return ErrIDNotFound
// 	}
// 	return s.storage.Delete(id)
// }
