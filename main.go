package main

import (
	"fmt"
	"log"

	"github.com/mrsanchez02/go-database/pkg/product"
	"github.com/mrsanchez02/go-database/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)

}
