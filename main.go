package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/mrsanchez02/go-database/pkg/product"
	"github.com/mrsanchez02/go-database/pkg/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	// err := serviceProduct.Delete(3)

	// if err != nil {
	// 	log.Fatalf("Product.Update: %v", err)
	// }

	// fmt.Println(m)

	ms, err := serviceProduct.GetByID(3)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There's no product with the ID provided.")
	case err != nil:
		log.Fatalf("Product.GetByID: %v", err)
	default:
		fmt.Println(ms)
	}
}
