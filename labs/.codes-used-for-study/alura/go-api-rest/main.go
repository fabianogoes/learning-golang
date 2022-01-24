package main

import (
	"fmt"
	"github.com/fabianogoes/learning-golang/alura/go-api-rest/database"

	"github.com/fabianogoes/learning-golang/alura/go-api-rest/models"
	"github.com/fabianogoes/learning-golang/alura/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Personalidade 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Personalidade 22", Historia: "Historia 22"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
