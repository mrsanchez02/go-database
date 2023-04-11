package storage

import (
	"database/sql"
	"fmt"

	"github.com/mrsanchez02/go-database/pkg/invoiceitem"
)

const (
	MySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	MySQLCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES(?, ?)`
)

// MySQLInvoiceItem user to work with postgres - invoice_items
type MySQLInvoiceItem struct {
	db *sql.DB
}

// NewMySQLInvoiceItem return a new pointer of MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

// Migrate implement the interface invoiceitem.Storage
func (p *MySQLInvoiceItem) Migrate() error {
	smt, err := p.db.Prepare(MySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("InvoiceItem migration executed successfully")
	return nil
}

// CreateTx implements the interface invoiceHeader.Storage
func (p *MySQLInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(MySQLCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(headerID, item.ProductID)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		item.ID = uint(id)
	}
	return nil

}
