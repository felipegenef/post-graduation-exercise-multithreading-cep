# Consulta CEP API - Golang (Eng version below)

## Descrição

Este projeto faz parte de um exercício/exame de uma pós-graduação em Golang. Ele consiste em um script que consulta duas APIs externas para obter informações sobre o endereço de um CEP (Código de Endereçamento Postal) e retorna os dados mais rápidos obtidos. As APIs utilizadas são: **BrasilAPI** e **ViaCEP**. O script realiza requisições simultâneas para ambas as APIs e exibe a resposta mais rápida dentro de um tempo limite de 1 segundo.

## Funcionalidades

- Consulta duas APIs externas (BrasilAPI e ViaCEP) para obter informações de endereço a partir do CEP.
- Usa goroutines e canais (channels) para realizar as requisições simultaneamente.
- Exibe os dados do endereço a partir da API que responder mais rapidamente.
- Implementação de timeout para garantir que, caso as APIs não respondam dentro de 1 segundo, o script informe o erro de timeout.

## Requisitos

- Go 1.23.3 ou superior
- Acesso à internet para consumir as APIs externas (https://brasilapi.com.br/api/cep e http://viacep.com.br/ws/)

## Como Rodar o Script

### 1. Clone o Repositório

```bash
git clone https://github.com/felipegenef/post-graduation-exercise-multithreading-cep.git
cd consulta-cep-api
```

### 2. Instale as Dependências

Este projeto não possui dependências externas além da biblioteca padrão do Go, então você pode pular esta etapa.

### 3. Execute o Script

Para executar o script, use o seguinte comando:

```bash
go run main.go
```

Isso executará o script, que fará as requisições para as APIs e exibirá o endereço obtido da API que responder mais rapidamente.

### 4. Como Testar

Para testar o script, basta alterar o valor da variável cep na função main() para o CEP desejado, por exemplo:

```go
cep := "01153000"
```

Após rodar o script, ele fará as requisições para as APIs e exibirá os dados do endereço, como o logradouro, bairro, cidade e estado.

### 5. Exemplo de Saída

A saída será similar a:

```txt
API: ViaCEP
CEP: 05086-010
Logradouro: Rua João Afonso
Bairro: Vila Hamburguesa
Cidade: São Paulo
UF: SP
```

Se ocorrer um erro, o sistema exibirá a mensagem correspondente, como no caso de timeout ou falha na consulta.

# Zip Code Lookup API - Golang

## Description

This project is part of an exercise/exam for a Golang postgraduate course. It consists of a script that queries two external APIs to retrieve address information from a ZIP code (CEP) and returns the fastest response. The APIs used are BrasilAPI and ViaCEP. The script makes simultaneous requests to both APIs and displays the fastest response within a 1-second timeout.

## Features

- Queries two external APIs (BrasilAPI and ViaCEP) to obtain address information from a ZIP code.
- Uses goroutines and channels to make requests simultaneously.
- Displays the address data from the fastest responding API.
- Implements a timeout to ensure that, if the APIs do not respond within 1 second, the script reports a timeout error.

## Requirements

- Go 1.23.3 or higher
- Internet access to query the external APIs (https://brasilapi.com.br/api/cep and http://viacep.com.br/ws/)

## How to Run the Script

### 1. Clone the Repository

```bash
git clone https://github.com/felipegenef/post-graduation-exercise-multithreading-cep.git
cd consulta-cep-api
```

### 2. Install Dependencies

This project does not have external dependencies beyond the Go standard library, so you can skip this step.

### 3. Run the Script

To run the script, use the following command:

```bash
go run main.go
```

This will execute the script, which will query the APIs and display the address data from the fastest response.

### 4. How to Test

To test the script, simply change the value of the cep variable in the main() function to the desired ZIP code, for example:

```go
cep := "01153000"
```
After running the script, it will query the APIs and display the address data such as the street, neighborhood, city, and state.

### 5. Example Output

The output will look similar to this:

```txt
API: ViaCEP
CEP: 05086-010
Street: Rua João Afonso
Neighborhood: Vila Hamburguesa
City: São Paulo
State: SP
```

If an error occurs, the system will display the corresponding message, such as in the case of a timeout or failure in querying the API.
