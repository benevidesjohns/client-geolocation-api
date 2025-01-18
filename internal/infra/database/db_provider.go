package database

import (
	"database/sql"

	interfaces "github.com/benevidesjohns/client-geolocation-api/internal/app/repository"
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

// Função que cria uma nova instância do DBProvider com as dependências necessárias
func NewDBProvider(db *sql.DB) (*DBProvider, error) {
	// Inicializa o repositório de cliente, que depende do banco conectado
	clientRepo := repository.NewClientRepository(db)

	return &DBProvider{
		ClientRepo: clientRepo,
	}, nil
}
