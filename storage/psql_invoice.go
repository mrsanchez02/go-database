package storage

import (
	"database/sql"
	"fmt"

	"github.com/mrsanchez02/go-database/pkg/invoice"
	"github.com/mrsanchez02/go-database/pkg/invoiceheader"
	"github.com/mrsanchez02/go-database/pkg/invoiceitem"
)

// PsqlInvoice used for work with postgres - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewPsqlInvoice return a new pointer of PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("header: %w", err)
	}
	fmt.Printf("Invoice created with id: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("items: %w", err)
	}
	fmt.Printf("Items created: %d \n", len(m.Items))

	return tx.Commit()

}
