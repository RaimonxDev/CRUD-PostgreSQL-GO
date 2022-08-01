package main

import (
	"fmt"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/product"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/storage"
	"log"
)

func main() {
	storage.NewPSQLDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()

	if err != nil {
		log.Fatalf("Product.GetAll: %v", err)
	}

	fmt.Println(ms)

}
