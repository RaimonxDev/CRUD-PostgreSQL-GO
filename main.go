package main

import (
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/product"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/storage"
	"log"
)

func main() {
	storage.NewPSQLDB()
	//storage.Pool() nos devuelve la instancia de psql.DB
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

}
