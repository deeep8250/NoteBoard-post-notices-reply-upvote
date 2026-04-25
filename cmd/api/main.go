package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	handler "github.com/threadpulse/internal/auth/handlers"
	"github.com/threadpulse/internal/auth/repository"
	service "github.com/threadpulse/internal/auth/services"

	"github.com/threadpulse/internal/routes"
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

	for range 5 {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("error while connecting to db : ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed")

	}

	fmt.Println("database connect successfully")

	AuthRepo := repository.NewAuthRepo(db)
	AuthService := service.NewAuthService(AuthRepo)
	AuthHandler := handler.NewAuthHandler(AuthService)

	r := gin.Default()
	routes.Routes(r, AuthHandler)

	r.Run(":8080")

}
