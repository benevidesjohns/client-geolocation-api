package usecase

import (
	"fmt"
	"time"

	"github.com/benevidesjohns/client-geolocation-api/internal/app/domain/models"
	"github.com/benevidesjohns/client-geolocation-api/internal/app/repository"
)

/*
ClientUseCase define o contrato para as operações de negócio relacionadas a clientes.
Esta interface segue o princípio de segregação de interfaces (ISP) e permite
que diferentes implementações sejam facilmente trocadas ou mockadas para testes.
*/
type ClientUseCase interface {
	CreateClient(client *models.Client) error
	GetAllClients() ([]*models.Client, error)
	GetClientByID(id int) (*models.Client, error)
	GetClientsByCity(city string) ([]*models.Client, error)
	UpdateClient(client *models.Client) error
	DeleteClient(id int) error
	DeleteAllClients() error
}

/*
Estrutura que implementa a interface ClientUseCase.
Esta estrutura serve como uma camada intermediária entre a lógica de negócio
e o repositório, garantindo que as regras sejam aplicadas corretamente.
*/
type clientUseCase struct {
	clientRepo repository.ClientRepository
}

// NewClientUseCase cria uma nova instância do caso de uso para clientes
func NewClientUseCase(repo repository.ClientRepository) ClientUseCase {
	return &clientUseCase{
		clientRepo: repo,
	}
}

// CreateClient cria um novo cliente após aplicar regras de negócio, como validações
func (uc *clientUseCase) CreateClient(client *models.Client) error {
	if client.Name == "" {
		return fmt.Errorf("client name cannot be empty")
	}
	if client.City == "" || client.Country == "" {
		return fmt.Errorf("client city and country cannot be empty")
	}

	client.CreatedAt = time.Now()
	client.UpdatedAt = time.Now()
	err := uc.clientRepo.Create(client)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	return nil
}

// GetAllClients retorna todos os clientes cadastrados no sistema
func (uc *clientUseCase) GetAllClients() ([]*models.Client, error) {
	clients, err := uc.clientRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clients: %w", err)
	}
	return clients, nil
}

// GetClientByID busca um cliente pelo ID e aplica regras de validação
func (uc *clientUseCase) GetClientByID(id int) (*models.Client, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid client ID")
	}

	client, err := uc.clientRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve client by ID: %w", err)
	}

	return client, nil
}

// GetClientsByCity retorna todos os clientes de uma cidade específica
func (uc *clientUseCase) GetClientsByCity(city string) ([]*models.Client, error) {
	if city == "" {
		return nil, fmt.Errorf("city cannot be empty")
	}

	clients, err := uc.clientRepo.GetByCity(city)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clients by city: %w", err)
	}

	return clients, nil
}

// UpdateClient atualiza as informações de um cliente no sistema
func (uc *clientUseCase) UpdateClient(client *models.Client) error {
	if client.ID <= 0 {
		return fmt.Errorf("invalid client ID")
	}
	if client.Name == "" {
		return fmt.Errorf("client name cannot be empty")
	}

	client.UpdatedAt = time.Now()
	err := uc.clientRepo.Update(client)
	if err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}

	return nil
}

// DeleteClient exclui um cliente do sistema pelo ID
func (uc *clientUseCase) DeleteClient(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid client ID")
	}

	err := uc.clientRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete client: %w", err)
	}

	return nil
}

// DeleteAllClients exclui todos os clientes do sistema
func (uc *clientUseCase) DeleteAllClients() error {
	err := uc.clientRepo.DeleteAll()
	if err != nil {
		return fmt.Errorf("failed to delete all clients: %w", err)
	}

	return nil
}
