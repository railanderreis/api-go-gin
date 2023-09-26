package main

import (
	"github.com/railanderreis/api-go-gin/database"
	"github.com/railanderreis/api-go-gin/models"
	"github.com/railanderreis/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "reis 1", CPF: "123", RG: "123"},
		{Nome: "reis 2", CPF: "123", RG: "123"},
		{Nome: "reis 3", CPF: "123", RG: "123"},
	}
	routes.HandleRequests()
}
