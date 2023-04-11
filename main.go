package main

import (
	"log"

	"github.com/mrsanchez02/go-database/pkg/product"
	"github.com/mrsanchez02/go-database/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(3)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}

}
