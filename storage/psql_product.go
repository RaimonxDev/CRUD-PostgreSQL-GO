package storage

import (
	"database/sql"
	"fmt"
	"github.com/RaimonxDev/CRUD-PostgreSQL-GO/pkg/product"
)

// Psql product usado para trabajar con postgress
type PsqlProduct struct {
	db *sql.DB
}

/*Interfaces Propia para manejar el scan, como Query.rows tambien aplica Scan, Entonces cumple con nuestra interfaces y
podemos aplicarla en nuestra funciona helper ScanRow Products*/
type scanner interface {
	Scan(dest ...interface{}) error
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
	createProduct = `INSERT INTO products(name, observation, price, created_at) VALUES($1 , $2, $3, $4) RETURNING id`

	getAllProduct = `SELECT id, name, observation, price, created_at, updated_at FROM products`
	// reusamos la consulta de getAllProduct mas el where
	getProductById = getAllProduct + " WHERE id = $1"
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

// Create Implementa interface storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(createProduct)

	if err != nil {
		return err
	}
	defer stmt.Close()

	// RECIBE TODOS LOS PARAMETROS EN EL MISMO ORDEN DE LA CONSULTA

	// Usamos el helper para manejar los nulos en PSQL
	err = stmt.QueryRow(m.Name, stringToNull(m.Observation), m.Price, m.CreatedAt).Scan(&m.ID)

	if err != nil {
		return err
	}
	fmt.Println("Se creo el producto correctamente")
	return nil

}

func (p *PsqlProduct) GetAll() (product.Models, error) {

	stmt, err := p.db.Prepare(getAllProduct)
	if err != nil {
		//retornamos nil porque GET ALL retorna una product.Models y un error
		//Entonces para el slices usamos un nil
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//creamos un slices de product model

	ms := make(product.Models, 0)
	for rows.Next() {
		// rows implementa Scan y cumple con nuestra interface scanner y por eso enviamos row a nuestro scanRowProducts
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil

}

func (p *PsqlProduct) GeyByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(getProductById)

	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()
	/*Retonarnamos directamente porque scanRowProduct retonar lo mismo que getByID*/
	return scanRowProduct(stmt.QueryRow(id))
}

// Creamos esta funcion helper para que podamos usar en los metodos de getAll y getById
// Porque se repito el codigo en ambos methodos
/*Creamos una interface llamada scanner que implemente Scan,
y asi poder accerder a sus valores
*/
func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull)
	if err != nil {
		return &product.Model{}, err
	}
	m.Observation = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
