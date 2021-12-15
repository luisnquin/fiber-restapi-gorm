package main

import (
	"log"

	conn "github.com/luisnquin/fiber-restapi-gorm/db"
	"github.com/luisnquin/fiber-restapi-gorm/routers"
)

func main() {
	conn.Init()
	defer conn.DB.Close()

	app := routers.Init()
	log.Fatal(app.Listen(":8000"))
}
