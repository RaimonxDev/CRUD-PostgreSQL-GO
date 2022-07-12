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

func LoadEnvDB() map[string]string {
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

	envDB := LoadEnvDB()
	once.Do(func() {

		var err error
		db, err = sql.Open("postgres", envDB["database"]+"://"+envDB["userDB"]+":"+envDB["passwordDB"]+"@"+envDB["dbHost"]+"/"+envDB["databaseName"]+"?sslmode=disable")

		if err != nil {
			log.Fatalf("No se pudo conectar a la bases de datos: %s ", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No hay respuesta de la bases de datos %v ", err)
		}

		fmt.Println("Contectado a PSQL")
	})
}

// retorna una unica instancia a multidb
func _() *sql.DB {
	return db
}
