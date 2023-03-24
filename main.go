package main

import (
	"fmt"
	"github.com/felipeazsantos/go-rest-api/database"
	"github.com/felipeazsantos/go-rest-api/models"
	"github.com/felipeazsantos/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")

	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
