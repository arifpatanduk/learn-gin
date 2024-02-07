package main

import (
	"learn-gin/database"
	"learn-gin/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}