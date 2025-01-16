package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/benevidesjohns/client-geolocation-api/internal/app/domain/models"
)

/*
Estrutura que implementa as operações de banco de dados para a entidade 'clients'.

Ela é responsável por encapsular a lógica de acesso e manipulação dos dados,
facilitando o agrupamento das operações comuns no banco
e a manutenção/reutilização de código.
*/
type ClientRepository struct {
	DB *sql.DB
}

/*
ALIAS para o tipo Client definido em models.

Isso permite usar o tipo Client diretamente nesse arquivo,
sem a necessidade de referenciar o pacote models a todo momento.
*/
type Client = models.Client

/*
Função que insere um novo cliente no banco de dados.

Recebe os dados de um cliente (do tipo Client) e executa a criação
no banco. Se ocorrer algum erro durante a criação, ele será retornado.
Caso contrário, será retornado nil, o que indica sucesso na criação.
*/
func (repo *ClientRepository) Create(client *Client) error {
	query := `
		INSERT INTO clients (
			name, test, weight_kg, address, street, number, neighborhood,
			complement, city, state, country, latitude, longitude, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := repo.DB.Exec(query,
		client.Name, client.Test, client.WeightKG, client.Address,
		client.Street, client.Number, client.Neighborhood, client.Complement,
		client.City, client.State, client.Country, client.Latitude, client.Longitude,
		client.CreatedAt, client.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("error to insert client: %v", err)
	}

	log.Println("client inserted successfully")

	return nil
}

/*
Função que obtém todos os registros de clientes no banco de dados.

Realiza uma consulta ao banco e retorna uma lista de ponteiros para
os registros (do tipo Client). Se houver algum erro na consulta,
ele será retornado. Caso contrário, retorna a lista de clientes.
*/
func (repo *ClientRepository) GetAll() ([]*Client, error) {
	query := `SELECT * FROM clients`

	res, err := repo.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying clients: %v", err)
	}
	defer res.Close()

	var clients = []*Client{}

	for res.Next() {
		var client Client

		err := res.Scan(
			&client.ID, &client.Name, &client.Test, &client.WeightKG,
			&client.Address, &client.Street, &client.Number, &client.Neighborhood,
			&client.Complement, &client.City, &client.State, &client.Country,
			&client.Latitude, &client.Longitude, &client.CreatedAt, &client.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		clients = append(clients, &client)
	}

	return clients, nil
}

/*
Função que obtém um cliente específico pelo ID no banco de dados.

Recebe o ID como parâmetro e retorna o registro cliente que tenha esse ID,
se encontrado. Caso o cliente não seja encontrado, um erro será retornado.
Se houver outro erro durante a consulta, ele também será retornado.
*/
func (repo *ClientRepository) GetByID(id int) (*Client, error) {
	query := `
		SELECT id, name, test, weight_kg, address, street, number, neighborhood,
		       complement, city, state, country, latitude, longitude, created_at, updated_at
		FROM clients 
		WHERE id = ?
	`

	var client Client
	err := repo.DB.QueryRow(query, id).Scan(
		&client.ID, &client.Name, &client.Test, &client.WeightKG,
		&client.Address, &client.Street, &client.Number, &client.Neighborhood,
		&client.Complement, &client.City, &client.State, &client.Country,
		&client.Latitude, &client.Longitude, &client.CreatedAt, &client.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("client with ID %d not found", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error querying client by ID: %v", err)
	}

	return &client, nil
}

/*
Função que obtém todos os clientes de uma cidade específica no banco de dados.

Recebe o nome da cidade como parâmetro e retorna uma lista de clientes
que pertencem a essa cidade. Se ocorrer algum erro durante a consulta,
ele será retornado. Caso contrário, retorna a lista de clientes.
*/
func (repo *ClientRepository) GetByCity(city string) ([]*Client, error) {
	query := `SELECT * FROM clients WHERE city = ?`

	res, err := repo.DB.Query(query, city)
	if err != nil {
		return nil, fmt.Errorf("error to get fetch clients: %v", err)
	}
	defer res.Close()

	var clients []*Client

	for res.Next() {
		var client Client

		err := res.Scan(
			&client.ID, &client.Name, &client.Test, &client.WeightKG,
			&client.Address, &client.Street, &client.Number, &client.Neighborhood,
			&client.Complement, &client.City, &client.State, &client.Country,
			&client.Latitude, &client.Longitude, &client.CreatedAt, &client.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		clients = append(clients, &client)
	}

	return clients, nil
}

/*
Função que atualiza os dados de um cliente no banco de dados.

Recebe os dados atualizados do cliente (do tipo Client).
Retorna um erro caso algo dê errado durante a atualização ou se o cliente
não for encontrado. Caso contrário, a atualização é realizada com sucesso.
*/
func (repo *ClientRepository) Update(client *Client) error {
	query := `
		UPDATE clients
		SET name = ?, test = ?, weight_kg = ?, address = ?, street = ?, number = ?, neighborhood = ?, 
			complement = ?, city = ?, state = ?, country = ?, latitude = ?, longitude = ?, updated_at = NOW()
		WHERE id = ?
	`

	res, err := repo.DB.Exec(query,
		client.Name, client.Test, client.WeightKG, client.Address,
		client.Street, client.Number, client.Neighborhood, client.Complement,
		client.City, client.State, client.Country, client.Latitude, client.Longitude,
		client.ID,
	)
	if err != nil {
		return fmt.Errorf("error to update client: %v", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error to check update: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("client with ID %d not found", client.ID)
	}

	log.Println("client updated successfully")

	return nil
}

/*
Função que exclui um registro de cliente do banco de dados com base no ID fornecido.

Recebe o ID do cliente como parâmetro e, caso o cliente seja encontrado,
ele será deletado. Se ocorrer algum erro durante a exclusão ou se o cliente
não for encontrado, o erro será retornado. Caso contrário, a exclusão é
realizada com sucesso.
*/
func (repo *ClientRepository) Delete(id int) error {
	query := `
		DELETE FROM clients WHERE id = ?
	`

	res, err := repo.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error to delete client: %v", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking deletion: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("client with ID %d not found", id)
	}

	log.Println("client deleted successfully")

	return nil
}

/*
Função que exclui todos os registros de clientes do banco de dados.

Caso ocorra algum erro durante a exclusão, ele será retornado.
Se nenhum registro for excluído, é retornado um erro indicando isso.
Caso contrário, todos os registros são excluídos com sucesso.
*/
func (repo *ClientRepository) DeleteAll() error {
	query := `DELETE FROM clients`

	res, err := repo.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error deleting all clients: %v", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking deletion: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("no clients found to delete")
	}

	log.Printf("%d clients deleted successfully\n", ra)
	return nil
}
