package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
	// "net/http"
)

var (
	db *sql.DB
)

// Client -> clients table in database
type Client struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Test         string    `json:"test"`
	WeightKG     float64   `json:"weight_kg"`
	Address      string    `json:"address"`
	Street       string    `json:"street"`
	Number       string    `json:"number"`
	Neighborhood string    `json:"neighborhood"`
	Complement   string    `json:"complement"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Country      string    `json:"country"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// type handleTest struct{}

// func (h handleTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "test")
// }

func main() {
	// h := handleTest{}

	// http.Handle("/", h)

	// log.Println("server starting... port: 8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	var err error

	db, err = sql.Open("mysql", "user:password@tcp(172.19.0.2:3306)/geo-clients-db?parseTime=true")
	if err != nil {
		panic(err)
	}

	newClient := Client{
		Name:         "Joao Benevides",
		Test:         "Teste",
		WeightKG:     77.00,
		Address:      "Rua Teste, 123",
		Street:       "Rua Teste",
		Number:       "123",
		Neighborhood: "Centro",
		Complement:   "Apartamento 10",
		City:         "Jequié",
		State:        "BA",
		Country:      "Brasil",
		Latitude:     -23.550520,
		Longitude:    -46.633308,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = createClient(newClient)
	if err != nil {
		log.Fatalf("Error to insert client: %v\n", err)
	}

	clients, err := getAllClients()
	if err != nil {
		log.Fatalf("Error to get all clients: %v\n", err)
	}

	log.Printf("\nGet all clients")
	for _, c := range clients {
		log.Printf("Client: %v", c)
	}
	log.Printf("\n")

	city := "Bahia"
	clients, err = getClientsByCity(city)
	if err != nil {
		log.Fatalf("Error to get all clients by city: %v", err)
	}

	log.Printf("\nGet clients by city")
	for _, c := range clients {
		fmt.Printf("Client: %v", c)
	}

	updatedClient := Client{
		ID:           2,
		Name:         "Joao silva",
		Test:         "Teste 2",
		WeightKG:     75,
		Address:      "Rua teste 2, 456",
		Street:       "Rua teste 2",
		Number:       "456",
		Neighborhood: "Bairro teste",
		Complement:   "Apto 101",
		City:         "Belford Roxo",
		State:        "Rio de Janeiro",
		Country:      "Brasil",
		Latitude:     40.7128,
		Longitude:    -74.0060,
	}

	err = updateClient(updatedClient)
	if err != nil {
		log.Fatalf("Error to update client: %v", err)
	}

	client, err := getClientByID(1)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nClient by ID (1): %v", client)

	err = deleteClient(1)
	if err != nil {
		log.Fatalf("Error to delete client: %v\n", err)
	}
}

func createClient(client Client) error {
	query := `
		INSERT INTO clients (
			name, test, weight_kg, address, street, number, neighborhood,
			complement, city, state, country, latitude, longitude, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
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

func getAllClients() ([]*Client, error) {
	query := `SELECT * FROM clients`

	res, err := db.Query(query)
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

func getClientByID(id int) (*Client, error) {
	query := `
		SELECT id, name, test, weight_kg, address, street, number, neighborhood,
		       complement, city, state, country, latitude, longitude, created_at, updated_at
		FROM clients 
		WHERE id = ?
	`

	var client Client
	err := db.QueryRow(query, id).Scan(
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

func getClientsByCity(city string) ([]*Client, error) {
	query := `SELECT * FROM clients WHERE city = ?`

	res, err := db.Query(query, city)
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

func updateClient(client Client) error {
	query := `
		UPDATE clients
		SET name = ?, test = ?, weight_kg = ?, address = ?, street = ?, number = ?, neighborhood = ?, 
			complement = ?, city = ?, state = ?, country = ?, latitude = ?, longitude = ?, updated_at = NOW()
		WHERE id = ?
	`

	res, err := db.Exec(query,
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

func deleteClient(id int) error {
	query := `
		DELETE FROM clients WHERE id = ?
	`

	res, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error to delete client: %v", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error to check deletion: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("client with ID %d not found", id)
	}

	log.Println("client deleted successfully")

	return nil
}
