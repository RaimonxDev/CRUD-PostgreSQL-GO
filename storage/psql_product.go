package storage

import (
	"database/sql"
	"fmt"
)

// Psql product usado para trabajar con postgress
type PsqlProduct struct {
	db *sql.DB
}

const (
	//migrateProduct
	migrateProduct = `CREATE TABLE IF NOT EXISTS products(
	id SERIAL NOT NULL,
	name VARCHAR(100) NOT NULL,
	observation VARCHAR(150),
	price INT NOT NULL,
	created_At TIMESTAMP NOT NULL DEFAULT now(),
	updated_At TIMESTAMP,
	CONSTRAINT product_id_pk PRIMARY KEY(id)
)`
)

// Retorna un nuevo punto de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Metodo Migrate de interface Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(migrateProduct)
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
	fmt.Println("Migracion creada correctamente")
	return nil
}
