package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Structs para as respostas das APIs
// Struct to hold the response from ViaCEP API
type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

// Struct para a resposta da BrasilAPI
// Struct to hold the response from BrasilAPI
type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

// Struct para a resposta final que contém o endereço e a fonte
// Struct for the final response that holds the address and the source
type ApiResponse struct {
	Address interface{}
	Source  string
}

// Função para buscar os dados usando a API BrasilAPI
// Function to fetch data using the BrasilAPI
func fetchFromBrasilAPI(cep string, ch chan ApiResponse) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		// Caso haja erro na requisição, envia resposta com erro
		// If there is an error in the request, send an error response
		ch <- ApiResponse{Source: "BrasilAPI", Address: nil}
		return
	}
	defer resp.Body.Close()

	// Verifica se a resposta foi OK (status 200)
	// Check if the response status is OK (status 200)
	if resp.StatusCode != http.StatusOK {
		ch <- ApiResponse{Source: "BrasilAPI", Address: nil}
		return
	}

	var address BrasilAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		// Se falhar ao decodificar os dados, envia resposta com erro
		// If decoding the data fails, send an error response
		ch <- ApiResponse{Source: "BrasilAPI", Address: nil}
		return
	}

	// Envia os dados recebidos pela API
	// Send the data received from the API
	ch <- ApiResponse{Source: "BrasilAPI", Address: address}
}

// Função para buscar os dados usando a API ViaCEP
// Function to fetch data using the ViaCEP API
func fetchFromViaCEP(cep string, ch chan ApiResponse) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json", cep)
	resp, err := http.Get(url)
	if err != nil {
		// Caso haja erro na requisição, envia resposta com erro
		// If there is an error in the request, send an error response
		ch <- ApiResponse{Source: "ViaCEP", Address: nil}
		return
	}
	defer resp.Body.Close()

	// Verifica se a resposta foi OK (status 200)
	// Check if the response status is OK (status 200)
	if resp.StatusCode != http.StatusOK {
		ch <- ApiResponse{Source: "ViaCEP", Address: nil}
		return
	}

	var address ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		// Se falhar ao decodificar os dados, envia resposta com erro
		// If decoding the data fails, send an error response
		ch <- ApiResponse{Source: "ViaCEP", Address: nil}
		return
	}

	// Envia os dados recebidos pela API
	// Send the data received from the API
	ch <- ApiResponse{Source: "ViaCEP", Address: address}
}

func main() {
	// Defina o CEP desejado
	// Define the desired ZIP code
	cep := "01153000"

	// Canais para as respostas das duas APIs
	// Channels to receive responses from both APIs
	chBrasilAPI := make(chan ApiResponse)
	chViaCEP := make(chan ApiResponse)

	// Timeout para a resposta (1 segundo)
	// Timeout for the response (1 second)
	timeout := time.After(1 * time.Second)

	// Inicia as goroutines para as duas APIs
	// Start the goroutines for both APIs
	go fetchFromBrasilAPI(cep, chBrasilAPI)
	go fetchFromViaCEP(cep, chViaCEP)

	// Usando select para esperar a resposta mais rápida
	// Using select to wait for the fastest response
	select {
	case res := <-chBrasilAPI:
		if res.Address != nil {
			// Exibe os dados recebidos da API BrasilAPI
			// Display the data received from the BrasilAPI
			fmt.Printf("API: %s\n", res.Source)
			address := res.Address.(BrasilAPIResponse)
			fmt.Printf("CEP: %s\n", address.CEP)
			fmt.Printf("Logradouro: %s\n", address.Street)
			fmt.Printf("Bairro: %s\n", address.Neighborhood)
			fmt.Printf("Cidade: %s\n", address.City)
			fmt.Printf("UF: %s\n", address.State)
		} else {
			// Caso a resposta da BrasilAPI esteja vazia
			// If the response from BrasilAPI is empty
			fmt.Println("Erro: Não foi possível obter os dados do endereço da BrasilAPI.")
		}
	case res := <-chViaCEP:
		if res.Address != nil {
			// Exibe os dados recebidos da API ViaCEP
			// Display the data received from the ViaCEP API
			fmt.Printf("API: %s\n", res.Source)
			address := res.Address.(ViaCEPResponse)
			fmt.Printf("CEP: %s\n", address.CEP)
			fmt.Printf("Logradouro: %s\n", address.Logradouro)
			fmt.Printf("Bairro: %s\n", address.Bairro)
			fmt.Printf("Cidade: %s\n", address.Localidade)
			fmt.Printf("UF: %s\n", address.UF)
		} else {
			// Caso a resposta da ViaCEP esteja vazia
			// If the response from ViaCEP is empty
			fmt.Println("Erro: Não foi possível obter os dados do endereço da ViaCEP.")
		}
	case <-timeout:
		// Caso o tempo de resposta tenha excedido 1 segundo
		// If the response time exceeds 1 second
		fmt.Println("Erro: Timeout. Nenhuma resposta recebida dentro do tempo limite.")
	}
}
