package storage

import (
	"database/sql"
	"fmt"

	"github.com/mrsanchez02/go-database/pkg/product"
)

const (
	MySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	MySQLCreateProduct  = `INSERT INTO products(name, observations, price, created_at) VALUES (?, ?, ?, ?)`
	MySQLGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	MySQLGetProductByID = MySQLGetAllProduct + " WHERE id = ?"
)

// MySQLProduct user to work with postgres - product
type MySQLProduct struct {
	db *sql.DB
}

// NewMySQLProduct return a new pointer of MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate implement the interface product.Storage
func (p *MySQLProduct) Migrate() error {
	smt, err := p.db.Prepare(MySQLMigrateProduct)
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
func (p *MySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(MySQLGetAllProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreateAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)

	fmt.Printf("Product Id %v created successfully\n", m.ID)
	return nil
}

// GetAll implements interface for product.Storage
func (p *MySQLProduct) GetAll() (product.Models, error) {
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
func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
	smt, err := p.db.Prepare(MySQLGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}

	defer smt.Close()

	return scanRowProduct(smt.QueryRow(id))
}
