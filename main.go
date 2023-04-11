package main

import (
	"fmt"
	"log"

	"github.com/mrsanchez02/go-database/pkg/product"
	"github.com/mrsanchez02/go-database/storage"
)

func main() {
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
}
