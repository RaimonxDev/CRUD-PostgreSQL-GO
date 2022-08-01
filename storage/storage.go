package storage

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

// Global Variables
var (
	db   *sql.DB
	once sync.Once
)

func loadEnvDB() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar .env %v", err)
	}

	envDB := map[string]string{
		"database":     os.Getenv("DATABASE"),
		"dbHost":       os.Getenv("DATABASE_HOST"),
		"userDB":       os.Getenv("DATABASE_USER"),
		"passwordDB":   os.Getenv("DATABASE_PASSWORD"),
		"databaseName": os.Getenv("DATABASE_NAME"),
		"databasePort": os.Getenv("DATABASE_PORT"),
	}

	return envDB
}

// Patron Singleton nos ayudamos de once para ejecutar la conexion de la bases de datos 1 sola vez
func NewPSQLDB() {

	envDB := loadEnvDB()
	once.Do(func() {
		// manejar solo los errores con una variable unica llamada err
		var err error
		// db hace referencia a la variable global
		db, err = sql.Open("postgres", envDB["database"]+"://"+envDB["userDB"]+":"+envDB["passwordDB"]+"@"+envDB["dbHost"]+"/"+envDB["databaseName"]+"?sslmode=disable")

		if err != nil {
			log.Fatalf("No se pudo conectar a la bases de datos: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No hay respuesta de la bases de datos %v ", err)
		}

		fmt.Println("Contectado a PSQL")
	})
}

// retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}

//funcion helper para controlar los parametro null
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	// Si null es distinto a vacio entonces hay algo que el usuario escribio
	if null.String != "" {
		null.Valid = true
	}
	// Si esta vacio entonces mandamos la variable null que hemos creado
	return null
}
