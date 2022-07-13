package storage

import (
	"database/sql"
	"fmt"
)

// Psql product usado para trabajar con postgress
type PsqlInvoiceHeader struct {
	db *sql.DB
}

const (
	//Migrate invoiceheader
	migrateInvoiceheader = `CREATE TABLE IF NOT EXISTS invoices_headers(
	id SERIAL NOT NULL,
	client VARCHAR(100) NOT NULL,
	created_At TIMESTAMP NOT NULL DEFAULT now(),
	updated_At TIMESTAMP,
	CONSTRAINT invoices_headers_id_pk PRIMARY KEY(id)
)`
)

// Retorna un nuevo punto de Invoiceheader

func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Method Migrate de interface Storage

func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(migrateInvoiceheader)
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
	fmt.Println("Migracion invoiceheader creada correctamente")
	return nil
}
