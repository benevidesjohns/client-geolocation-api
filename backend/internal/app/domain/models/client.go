package models

import (
	"time"
)

// Essa struct representa uma entidade cliente no banco de dados
// e é a representação do domínio de clientes que contém todos
// os campos necessários para georreferenciamento.
type Client struct {
	ID           int       `json:"id"`                               // Identificação do cliente
	Name         string    `json:"name" validate:"required"`         // Nome do cliente
	Test         string    `json:"test" validate:"required"`         // Campo de teste técnico
	WeightKG     float64   `json:"weight_kg" validate:"required"`    // Peso em kg
	Address      string    `json:"address" validate:"required"`      // Endereço completo para geocodificação
	Street       string    `json:"street" validate:"required"`       // Logradouro extraído da geocodificação
	Number       string    `json:"number" validate:"required"`       // Número do endereço
	Neighborhood string    `json:"neighborhood" validate:"required"` // Bairro
	Complement   string    `json:"complement" validate:"required"`   // Complemento do endereço
	City         string    `json:"city" validate:"required"`         // Cidade
	State        string    `json:"state" validate:"required"`        // Estado
	Country      string    `json:"country" validate:"required"`      // País
	Latitude     float64   `json:"latitude" validate:"required"`     // Latitude do endereço
	Longitude    float64   `json:"longitude" validate:"required"`    // Longitude do endereço
	CreatedAt    time.Time `json:"created_at"`                       // Data de criação do registro
	UpdatedAt    time.Time `json:"updated_at"`                       // Data de atualização do registro
}
