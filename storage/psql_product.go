package storage

import (
	"database/sql"
	"fmt"

	"github.com/mrsanchez02/go-database/pkg/product"
)

type Scanner interface {
	Scan(dest ...any) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct  = `INSERT INTO products(name, observations, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteProduct  = `DELETE FROM products WHERE id = $1`
)

// PsqlProduct user to work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	smt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product migration executed successfully")
	return nil
}

// Create implements interface for product.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreateAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("Product created successfully")
	return nil
}

// GetAll implements interface for product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetByID implement the interface product.Storage
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	smt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}

	defer smt.Close()

	return scanRowProduct(smt.QueryRow(id))
}

// Update implement the interface product.Storage
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %d doesn't exist ", m.ID)
	}

	fmt.Println("Product updated successfully")
	return nil
}

func (p *PsqlProduct) Delete(id uint) error {
	smt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer smt.Close()

	res, err := smt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %d doesn't exist ", id)
	}
	fmt.Println("Product deleted successfully")
	return nil

}

func scanRowProduct(s Scanner) (*product.Model, error) {
	m := &product.Model{}
	ObservationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&ObservationNull,
		&m.Price,
		&m.CreateAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = ObservationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
