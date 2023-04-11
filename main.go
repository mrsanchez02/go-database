package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/mrsanchez02/go-database/pkg/product"
	"github.com/mrsanchez02/go-database/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(2)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There's no product with this id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}

}
