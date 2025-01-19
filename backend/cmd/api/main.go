package main

import (
	"log"
	"net/http"

	"github.com/benevidesjohns/client-geolocation-api/internal/di"
)

func main() {
	// Inicializa o DIContainer com todas as dependências necessárias
	container, err := di.NewDIContainer()
	if err != nil {
		log.Fatal("Error to initialize DIContainer: ", err)
	}

	// Inicializa o servidor HTTP com as rotas configuradas
	log.Println("Server is running... Port: 8080")
	http.ListenAndServe(":8080", container.HTTPProvider.Router)
}
