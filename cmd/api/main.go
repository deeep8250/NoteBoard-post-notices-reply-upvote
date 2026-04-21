package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func main() {

	dsn := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal("Migration files not found ", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migtaion failed ", err.Error())
	}

	log.Println("Migrations ran successfully")

	var db *sqlx.DB

	// for range 5 {
	// 	db, err = sqlx.Connect("postgres", dsn)
	// }
	db, err = sqlx.Connect("postgres", dsn)
	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed")

	}

	fmt.Println("database connect successfully")

}
