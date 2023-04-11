# Readme

## Products table migration

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
  log.Fatalf("Product.Migrate: %v", err)
}
```

## InvoiceHeaders table migration

```go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
  log.Fatalf("InvoiceHeader.Migrate: %v", err)
}
```

## InvoiceItems table migration

```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
  log.Fatalf("InvoiceItem.Migrate: %v", err)
}
```

## Create a product
```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m := &product.Model{
	Name:         "Curso de db con Go",
	Price:        70,
	Observations: "on fire",
}
if err := serviceProduct.Create(m); err != nil {
	log.Fatalf("product.Create: %v", err)
}

fmt.Printf("%+v\n", m)
```

## Create an invoice

```go
storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
storageInvoice := storage.NewPsqlInvoice(
	storage.Pool(),
	storageHeader,
	storageItems,
)

m := &invoice.Model{
	Header: &invoiceheader.Model{
		Client: "Eddy Abreu",
	},
	Items: invoiceitem.Models{
		&invoiceitem.Model{ProductID: 1},
	},
}

serviceInvoice := invoice.NewService(storageInvoice)
if err := serviceInvoice.Create(m); err != nil {
	log.Fatalf("invoice.Create: %v", err)
}
```