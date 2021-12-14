package main

import (
	"log"

	"github.com/luisnquin/fiber-restapi-gorm/db"
	"github.com/luisnquin/fiber-restapi-gorm/routers"
)

func main() {
	db.Init()

	app := routers.Init()
	log.Fatal(app.Listen(":8000"))
}
