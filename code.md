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
