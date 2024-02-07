package main

import (
	"Gin/database"
	"Gin/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()

	routers.StartServer().Run(PORT)
}
