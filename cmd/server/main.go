package main

import (
	"architecture/internal/route"
	"architecture/internal/storage/postgres"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db, err := postgres.Init("postgres://user:password@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	route.Init(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
