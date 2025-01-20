package main

import (
	"log"
	"net/http"

	"github.com/benevidesjohns/client-geolocation-api/internal/di"
	"github.com/rs/cors"
)

func main() {
	// Inicializa o DIContainer com todas as dependências necessárias
	container, err := di.NewDIContainer()
	if err != nil {
		log.Fatal("Error to initialize DIContainer: ", err)
	}

	// Configurações CORS
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		Debug:            true,
		AllowCredentials: true,
	}

	// Cria o middleware CORS
	corsHandler := cors.New(corsOptions).Handler(container.HTTPProvider.Router)

	// Inicializa o servidor HTTP com as rotas configuradas
	log.Println("Server is running... Port: 8080")
	http.ListenAndServe(":8080", corsHandler)
}
