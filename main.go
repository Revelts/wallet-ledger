package main

import (
	"go-testing/database"
	"go-testing/routes"
)

func main() {
	database.InitDB()
	routes.InitRoutes()
}
