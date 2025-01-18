package http

import (
	"github.com/benevidesjohns/client-geolocation-api/internal/app/usecase"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/http/handlers"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/http/routes"
	"github.com/gorilla/mux"
)

/*
Estrutura que armazena e gerencia as implementações dos handlers HTTP.
Esse HTTPProvider facilita a injeção de dependências, criando os handlers necessários
e permitindo que as dependências sejam passadas para os mesmos de maneira centralizada.
*/
type HTTPProvider struct {
	ClientHandler *handlers.ClientHandler
	Router        *mux.Router
}

// NewHTTPProvider cria uma nova instância do HTTPProvider com as dependências necessárias
func NewHTTPProvider(clientUseCase usecase.ClientUseCase) (*HTTPProvider, error) {
	// Cria o handler passando o caso de uso necessário
	clientHandler := handlers.NewClientHandler(clientUseCase)

	// Usa o arquivo router para configurar as rotas
	router := routes.SetupRoutes(clientHandler)

	return &HTTPProvider{
		ClientHandler: clientHandler,
		Router:        router,
	}, nil
}
