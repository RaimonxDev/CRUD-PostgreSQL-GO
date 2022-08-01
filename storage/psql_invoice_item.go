package storage

import (
	"database/sql"
	"fmt"
)

// Psql product usado para trabajar con postgress
type PsqlInvoiceItem struct {
	db *sql.DB
}

const (
	//Migrate invoice items
	migrateInvoiceitems = `CREATE TABLE IF NOT EXISTS invoice_items(
	id SERIAL NOT NULL,
	invoice_header_id INT NOT NULL,
	product_id INT NOT NULL,
	created_At TIMESTAMP NOT NULL DEFAULT now(),
	updated_At TIMESTAMP,
	CONSTRAINT invoice_items_id_pk PRIMARY KEY(id),
    CONSTRAINT invoice_items_invoices_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoices_headers (id) 
    ON UPDATE RESTRICT ON DELETE RESTRICT, 
    CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) 
    ON UPDATE RESTRICT ON DELETE RESTRICT
)`
)

// Retorna un nuevo punto de Invoiceheader

func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Method Migrate de interface Storage

func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(migrateInvoiceitems)
	if err != nil {
		return err
	}
	// INMEDIATAMENTE Debemos de cerrar la conexion para liberar recursos
	defer stmt.Close()

	// No necesitamos el resultado. Solo vamos a crear la tabla
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migracion invoice items creada correctamente")
	return nil
}
