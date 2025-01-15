package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
	"net/http"
)

var (
	db *sql.DB
)

// Client -> clients table in database
type Client struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Test         string       `json:"test"`
	WeightKG     float64      `json:"weight_kg"`
	Address      string       `json:"address"`
	Street       string       `json:"street"`
	Number       string       `json:"number"`
	Neighborhood string       `json:"neighborhood"`
	Complement   string       `json:"complement"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	Country      string       `json:"country"`
	Latitude     float64      `json:"latitude"`
	Longitude    float64      `json:"longitude"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type handleTest struct{}

func (h handleTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	var err error
	var insertError error

	db, err = sql.Open("mysql", "user:password@tcp(172.19.0.2:3306)/geo-clients-db?parseTime=true")
	if err != nil {
		panic(err)
	}

	client := Client{
		Name:         "Joao Benevides",
		Test:         "Teste",
		WeightKG:     77.00,
		Address:      "Rua Teste, 123",
		Street:       "Rua Teste",
		Number:       "123",
		Neighborhood: "Centro",
		Complement:   "Apartamento 10",
		City:         "Bahia",
		State:        "BA",
		Country:      "Brasil",
		Latitude:     -23.550520,
		Longitude:    -46.633308,
		CreatedAt:    time.Now(),
	}
	insertError = insertClient(client)
	if insertError != nil {
		log.Fatalf("Error to insert client: %v", insertError)
	}

	clients, err := getAllClients()
	if err != nil {
		log.Fatalf("Error to get all clients: %v", err)
	}

	for _, c := range clients {
		log.Printf("Client: %v", c)
	}

	h := handleTest{}

	http.Handle("/", h)

	log.Println("server starting... port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func insertClient(client Client) error {
	query := `
		INSERT INTO clients (
			name, test, weight_kg, address, street, number, neighborhood, 
			complement, city, state, country, latitude, longitude, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		client.Name, client.Test, client.WeightKG, client.Address,
		client.Street, client.Number, client.Neighborhood, client.Complement,
		client.City, client.State, client.Country, client.Latitude, client.Longitude,
		client.CreatedAt,
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
