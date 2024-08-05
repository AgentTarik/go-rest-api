package main

import (
	"fmt"
	"github.com/AgentTarik/go-rest-api/database"
	"github.com/AgentTarik/go-rest-api/models"
	"github.com/AgentTarik/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Bruno mock", Historia: "Homem Top"},
		{Id: 2, Nome: "Polna mock", Historia: "Homem do Fuzil"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Inicaindo...")
	routes.HandleRequest()
}
