package storage

import (
	"database/sql"
	"fmt"

	"github.com/mrsanchez02/go-database/pkg/invoiceheader"
)

const (
	MySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	MySQLCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)

// MySQLInvoiceHeader user to work with postgres - invoice_headers
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader return a new pointer of MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implement the interface invoiceHeader.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	smt, err := p.db.Prepare(MySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("InvoiceHeader migration executed successfully!")
	return nil
}

// CreateTx implements the interface invoiceHeader.Storage
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(MySQLCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)

	return nil
}
