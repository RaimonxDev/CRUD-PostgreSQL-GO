package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/product"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/storage"
	"log"
)

func main() {
	storage.NewPSQLDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(19)

	if err != nil {
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con ese ID")
	case err != nil:
		log.Fatalf("Product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}

}
