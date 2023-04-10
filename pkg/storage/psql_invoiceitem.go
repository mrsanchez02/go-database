package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY(product_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
)

// PsqInvoiceItem user to work with postgres - invoice_items
type PsqInvoiceItem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem return a new pointer of PsqInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqInvoiceItem {
	return &PsqInvoiceItem{db}
}

// Migrate implement the interface invoiceitem.Storage
func (p *PsqInvoiceItem) Migrate() error {
	smt, err := p.db.Prepare(psqlMigrateInvoiceItem)
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
