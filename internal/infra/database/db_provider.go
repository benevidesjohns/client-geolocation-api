package database

import (
	interfaces "github.com/benevidesjohns/client-geolocation-api/internal/app/repository"

	"github.com/benevidesjohns/client-geolocation-api/internal/infra/database/mysql"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/database/mysql/repository"
)

/*
Estrutura que armazena e gerencia as implementações do repositório (que seguem tal interface).

Ela é responsável por encapsular as dependências reacionadas ao banco de dados.
Esse DBProvider irá facilitar a injeção de dependências e os testes nessa aplicação.
*/
type DBProvider struct {
	ClientRepo interfaces.ClientRepository
}

func NewDBProvider() (*DBProvider, error) {
	// Inicializa conexão com o banco
	db, err := mysql.NewDBConnection()
	if err != nil {
		return nil, err
	}

	// Inicializa o repositório com a implementação concreta
	clientRepo := &repository.ClientRepository{DB: db}

	return &DBProvider{
		ClientRepo: clientRepo,
	}, nil
}
