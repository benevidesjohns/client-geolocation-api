package repository

import (
	"github.com/benevidesjohns/client-geolocation-api/internal/app/domain/models"
)

/*
ALIAS para o tipo Client definido em models.

Isso permite usar o tipo Client diretamente nesse arquivo,
sem a necessidade de referenciar o pacote models a todo momento.
*/
type Client = models.Client

// ClientRepository define o contrato para operações de persistência de clientes
// Essa interface segue o princípio de segregação de interfaces (ISP, um dos princípios do SOLID),
// definindo apenas os métodos necessários para o domínio de clientes.
type ClientRepository interface {
	// Função que deve adicionar um novo cliente no banco
	Create(client *Client) error

	// Função que deve recuperar todos os clientes do banco
	GetAll() ([]*Client, error)

	// Função que deve recuperar um cliente específico pelo seu ID
	GetByID(id int) (*Client, error)

	// Função que deve recuperar todos os clientes de uma determinada cidade
	GetByCity(city string) ([]*Client, error)

	// Função que deve atualizar os dados de um cliente existente
	Update(client *Client) error

	// Função que deve remover um cliente do banco
	Delete(id int) error

	// Função que deve remover todos os clientes do banco
	DeleteAll() error
}
