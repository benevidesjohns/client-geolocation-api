package routes

import (
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/http/handlers"
	"github.com/gorilla/mux"
)

// Função que cria e configura todas as rotas da API utilizando o Mux.
func SetupRoutes(clientHandler *handlers.ClientHandler) *mux.Router {
	router := mux.NewRouter()

	// Agrupa todas as rotas com o prefixo /api
	api := router.PathPrefix("/api").Subrouter()

	// Configuração das rotas de clientes
	clients := api.PathPrefix("/deliveries").Subrouter()

	// Rotas que não precisam do ID
	clients.HandleFunc("", clientHandler.CreateClient).Methods("POST")
	clients.HandleFunc("", clientHandler.GetAllClients).Methods("GET")
	clients.HandleFunc("", clientHandler.DeleteAllClients).Methods("DELETE")

	// Rotas que precisam do ID
	clients.HandleFunc("/{id:[0-9]+}", clientHandler.GetClientByID).Methods("GET")
	clients.HandleFunc("/{id:[0-9]+}", clientHandler.UpdateClient).Methods("PUT")
	clients.HandleFunc("/{id:[0-9]+}", clientHandler.DeleteClient).Methods("DELETE")

	return router
}
