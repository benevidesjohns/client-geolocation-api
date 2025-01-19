package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/benevidesjohns/client-geolocation-api/internal/app/domain/models"
	"github.com/benevidesjohns/client-geolocation-api/internal/app/usecase"
	"github.com/gorilla/mux"
)

/*
ClientHandler encapsula toda a lógica de manipulação de requisições HTTP relacionadas a clientes.
Essa estrutura segue o padrão de Clean Architecture, onde os handlers são a camada mais externa
e se comunicam com os casos de uso (usecase) que contêm a lógica de negócio.

A injeção de dependência é utilizada aqui pra fornecer o caso de uso necessário,
fazendo com que o código fique mais testável e flexível.
*/
type ClientHandler struct {
	clientUseCase usecase.ClientUseCase
}

// NewClientHandler cria uma nova instância de ClientHandler com as dependências necessárias.
func NewClientHandler(useCase usecase.ClientUseCase) *ClientHandler {
	return &ClientHandler{
		clientUseCase: useCase,
	}
}

/*
CreateClient manipula requisições POST para criar novos clientes.
Este handler:
1. Decodifica o JSON do corpo da requisição
2. Valida os dados recebidos
3. Chama o caso de uso apropriado para processar a lógica de negócio
4. Retorna a resposta adequada com o código de status HTTP correto
*/
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	// Decodifica o corpo da requisição JSON
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Println(client.WeightKG)

	// Chama o caso de uso para criar o cliente
	if err := h.clientUseCase.CreateClient(&client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna sucesso com código 201 Created
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Client created successfully",
	})
}

/*
GetAllClients manipula requisições GET para listar todos os clientes.
Suporta filtragem opcional por cidade através de query parameter.
*/
func (h *ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	// Verifica se há filtro por cidade
	city := r.URL.Query().Get("city")

	var clients []*models.Client
	var err error

	if city != "" {
		clients, err = h.clientUseCase.GetClientsByCity(city)
	} else {
		clients, err = h.clientUseCase.GetAllClients()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clients)
}

/*
GetClientByID manipula requisições GET para buscar um cliente específico por ID.
Extrai o ID da URL e valida antes de processar a requisição.
*/
func (h *ClientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	client, err := h.clientUseCase.GetClientByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(client)
}

/*
UpdateClient manipula requisições PUT para atualizar dados de um cliente existente.
Combina a validação do ID da URL com os dados do corpo da requisição.
*/
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client.ID = id
	if err := h.clientUseCase.UpdateClient(&client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Client updated successfully",
	})
}

/*
DeleteClient manipula requisições DELETE para remover um cliente específico.
Suporta tanto exclusão individual por ID quanto exclusão em massa.
*/
func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.clientUseCase.DeleteClient(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Client deleted successfully",
	})
}

/*
DeleteAllClients manipula requisições DELETE para remover todos os clientes.
Esta é uma operação potencialmente perigosa e deve ser usada com cautela.
*/
func (h *ClientHandler) DeleteAllClients(w http.ResponseWriter, r *http.Request) {
	if err := h.clientUseCase.DeleteAllClients(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "All clients deleted successfully",
	})
}
