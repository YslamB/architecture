package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/TurkmenistanRailways/payment/internal/config"
	"github.com/TurkmenistanRailways/payment/internal/route"
	"github.com/TurkmenistanRailways/payment/internal/storage/postgres"
)

func main() {
	conf := config.Init()
	db := postgres.Init()

	r := gin.Default()
	route.Init(r, db, conf)

	if err := r.Run(conf.PORT); err != nil {
		log.Fatal(err)
	}
}
