```go

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


create productos

m := &product.Model{
Name:        "Curso de go advanced POO",
Price:       900,
Observation: "Curso completo",
}
if err := serviceProduct.Create(m); err != nil {
log.Fatalf("Error al crear producto %v", err)

}

```