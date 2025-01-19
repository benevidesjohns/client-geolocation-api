package di

import (
	"fmt"

	"github.com/benevidesjohns/client-geolocation-api/internal/app/usecase"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/database"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/database/mysql"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/http"
)

// Estrutura que centraliza todos os providers da aplicação
type DIContainer struct {
	HTTPProvider *http.HTTPProvider
	DBProvider   *database.DBProvider
}

// Função cria e retorna um container com as dependências necessárias
func NewDIContainer() (*DIContainer, error) {
	// Inicializa conexão com o banco
	db, err := mysql.NewDBConnection()
	if err != nil {
		return nil, err
	}

	// Inicializa o DBProvider
	dbProvider, err := database.NewDBProvider(db)
	if err != nil {
		return nil, fmt.Errorf("error to initialize o DBProvider: %w", err)
	}

	// Inicializa o caso de uso de cliente, que depende do DBProvider
	clientUseCase := usecase.NewClientUseCase(dbProvider.ClientRepo)

	// Inicializa o HTTPProvider
	httpProvider, err := http.NewHTTPProvider(clientUseCase)
	if err != nil {
		return nil, fmt.Errorf("error to initialize HTTPProvider: %w", err)
	}

	return &DIContainer{
		HTTPProvider: httpProvider,
		DBProvider:   dbProvider,
	}, nil
}
