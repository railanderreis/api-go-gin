package main

import (
	"github.com/railanderreis/api-go-gin/database"
	"github.com/railanderreis/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
