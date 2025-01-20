// pkg/config/config.go

package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Estrutura que contém as configurações do banco de dados
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Função que carrega as configurações do arquivo .env
func LoadConfig() (*DBConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error to load .env file: %w", err)
	}

	config := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// Validação das configurações.
	// Define configurações padrões caso alguma configuração necessária
	// não seja encontrada no arquivo de variáveis de ambiente.
	if config.Host == "" || config.Port == "" || config.User == "" ||
		config.Password == "" || config.DBName == "" {
		config.Host = "localhost"
		config.Port = "3306"
		config.User = "user"
		config.Password = "password"
		config.DBName = "my-database"

		// return nil, fmt.Errorf("database configs not found")
	}

	return config, nil
}

/*
DSN retorna a string de conexão formatada

Motivo dos parâmetros utilizados:

-> parseTime=true, para habilitar o parse de colunas do tipo DATETIME ou TIMESTAMP para o time.Time;

-> charset-utf8, para definir o conjunto de caracteres usado pelo banco como UTF-8 e permitir caracteres especiais.

-> loc=America%2Sao_Paulo, para especificar o timezone usado nas operações com datas e horas.

-> tls==skip-verify, para configurar o driver para ignorar a verificação de certificados TLS/SSL
*/
func (c *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=America%%2FSao_Paulo&tls=skip-verify",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
