package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"log"

	"github.com/benevidesjohns/client-geolocation-api/internal/app/domain/models"
	"github.com/benevidesjohns/client-geolocation-api/internal/infra/database"
)

type Client = models.Client

func main() {
	dbProvider, err := database.NewDBProvider()
	if err != nil {
		log.Fatalf("error to initialize DB provider: %v", err)
	}

	clientRepo := dbProvider.ClientRepo

	// newClient := &Client{
	// 	Name:         "Joao Benevides",
	// 	Test:         "Teste",
	// 	WeightKG:     77.00,
	// 	Address:      "Rua Teste, 123",
	// 	Street:       "Rua Teste",
	// 	Number:       "123",
	// 	Neighborhood: "Centro",
	// 	Complement:   "Apartamento 10",
	// 	City:         "Jequié",
	// 	State:        "BA",
	// 	Country:      "Brasil",
	// 	Latitude:     -23.550520,
	// 	Longitude:    -46.633308,
	// 	CreatedAt:    time.Now(),
	// 	UpdatedAt:    time.Now(),
	// }

	// if err := clientRepo.Create(newClient); err != nil {
	// 	log.Fatalf("Erro ao criar cliente: %v", err)
	// }

	// clients, err := clientRepo.GetAll()
	// if err != nil {
	// 	log.Fatalf("Error to get all clients: %v\n", err)
	// }

	// log.Printf("\nGet all clients")
	// for _, c := range clients {
	// 	log.Printf("Client: %v", c)
	// }

	// city := "Jequié"
	// clients, err := clientRepo.GetByCity(city)
	// if err != nil {
	// 	log.Fatalf("Error to get all clients by city: %v", err)
	// }

	// log.Printf("\nGet clients by city")
	// for _, c := range clients {
	// 	fmt.Printf("Client: %v", c)
	// }

	// updatedClient := &Client{
	// 	ID:           2,
	// 	Name:         "Joao silva",
	// 	Test:         "Teste 2",
	// 	WeightKG:     75,
	// 	Address:      "Rua teste 2, 456",
	// 	Street:       "Rua teste 2",
	// 	Number:       "456",
	// 	Neighborhood: "Bairro teste",
	// 	Complement:   "Apto 101",
	// 	City:         "Belford Roxo",
	// 	State:        "Rio de Janeiro",
	// 	Country:      "Brasil",
	// 	Latitude:     40.7128,
	// 	Longitude:    -74.0060,
	// }

	// err = clientRepo.Update(updatedClient)
	// if err != nil {
	// 	log.Fatalf("Error to update client: %v", err)
	// }

	client, err := clientRepo.GetByID(2)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nClient by ID (2): %v", client)

	// err = clientRepo.Delete(1)
	// if err != nil {
	// 	log.Fatalf("Error to delete client: %v\n", err)
	// }

	// err = clientRepo.DeleteAll()
	// if err != nil {
	// 	log.Fatalf("Error to delete all client: %v\n", err)
	// }
}
