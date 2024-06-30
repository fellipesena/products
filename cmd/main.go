package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	runMigration()

	server := gin.Default()

	server.GET("/teste", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	})

	server.Run(":8080")
}

func runMigration(){
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/products?sslmode=disable")

	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"products", driver)

	if err != nil {
		panic(err)
	}

	m.Up()
}