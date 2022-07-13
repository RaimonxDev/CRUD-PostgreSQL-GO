package main

import (
	invoiceItem "github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/invoice-item"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/invoiceheader"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/product"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/storage"
	"log"
)

func main() {
	storage.NewPSQLDB()

	//	PRODUCTS
	//storage.Pool() nos devuelve la instancia de psql.DB
	storageProduct := storage.NewPsqlProduct(storage.Pool())

	//Service product necesita un objeto que implemente la interface storage
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
	// INVOICE HEADER
	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceHeader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiveheader migrate")
	}
	// INVOICE ITEM
	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceItem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("Invoice item migrate")
	}

}
