package main

import (
	"log"

	"github.com/luisnquin/fiber-restapi-gorm/routers"
	"github.com/luisnquin/fiber-restapi-gorm/db"
)


func main() {
	db.Init()
	
	app := routers.Init()
	log.Fatal(app.Listen(":8000"))
}