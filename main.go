package main

import (
	"log"

	"github.com/mrsanchez02/go-database/pkg/invoice"
	"github.com/mrsanchez02/go-database/pkg/invoiceheader"
	"github.com/mrsanchez02/go-database/pkg/invoiceitem"
	"github.com/mrsanchez02/go-database/storage"
)

func main() {
	storage.NewMySQLDB()

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewMySQLInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Jose Sanchez",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 2},
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}
