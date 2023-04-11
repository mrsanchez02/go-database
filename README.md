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

## Query all products

```go
storageProduct := storage.NewMySQLProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

ms, err := serviceProduct.GetAll()
if err != nil {
 log.Fatalf("product.GetAll: %v", err)
}

fmt.Println(ms)
```

## Update a Product

```go
storageProduct := storage.NewMySQLProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m := &product.Model{
  ID:    1,
  Name:  "Curso de CSS",
  Price: 200,
}
err := serviceProduct.Update(m)
if err != nil {
  log.Fatalf("product.Update: %v", err)
}
```

## Delete a Product

```go
storageProduct := storage.NewMySQLProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

err := serviceProduct.Delete(3)
if err != nil {
 log.Fatalf("product.Delete: %v", err)
}

```

## Create an invoice

```go
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
```

## DaoProduct

```go
driver := storage.MySQL
storage.New(driver)
myStorage, err := storage.DAOProduct(driver)
if err != nil {
  log.Fatalf("DAOProduct: %v", err)
}
serviceProduct := product.NewService(myStorage)
ms, err := serviceProduct.GetAll()
if err != nil {
  log.Fatalf("product.GetAll: %v", err)
}
fmt.Println(ms)
```
