package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/benevidesjohns/client-geolocation-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

/*
Função responsável por criar e gerenciar uma nova conexão com o banco de dados MySQL.

Essa função utiliza as configurações carregadas do arquivo .env
para construir a string de conexão e iniciar a conexão com o banco.

O motivo principal é abstrair todo o processo de configuração do banco,
fornecendo uma instância rápida de `*sql.DB` que pode ser usada pela aplicação.

Se algo der errado, como erro ao carregar as configurações
ou erro na conexão (por exemplo: erro no `Ping`),
a função retorna um erro para facilitar o tratamento.
*/
func NewDBConnection() (*sql.DB, error) {
	var dbConfig *configs.DBConfig

	// Carrega as configurações do banco de dados a partir do arquivo .env
	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("error to load configs: %v", err)
	}

	// Abre a conexão com o MySQL no formato necessário.
	db, err := sql.Open("mysql", dbConfig.DSN())
	if err != nil {
		return nil, fmt.Errorf("error to open mysql connection: %w", err)
	}

	// Verifica se a conexão está ativa
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error to connect to mysql: %w", err)
	}

	return db, nil
}
