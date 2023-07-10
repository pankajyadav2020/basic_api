package main

import (
	"basic_api/config"
	"basic_api/routes"
)

//const SecretKey = "secret"

func main() {
	config.Dbmigration()
	routes.InitRouter()
}
