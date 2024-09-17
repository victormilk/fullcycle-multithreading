# CEP Lookup Service

This project is a simple Go application that fetches address information based on a given CEP (Brazilian postal code) using two different APIs: BrasilAPI and ViaCepAPI. The application runs both API requests concurrently and returns the result from the API that responds first.

## Prerequisites

- Go
- Make

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/victormilk/fullcycle-multithreading.git
    cd fullcycle-multithreading
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

To run the application, you can use the provided Makefile. By default, it uses the CEP `01153000`.

1. Run the application:
    ```sh
    make run
    ```

2. To specify a different CEP:
    ```sh
    make run CEP=12345678
    ```

3. Using go cli
    ```sh
    go run main.go
    ```

4. To specify a different CEP using go cli:
    ```sh
    go run main.go 12345678
    ```

## Endpoints

- **BrasilAPI**: `https://brasilapi.com.br/api/cep/v1/{cep}`
- **ViaCepAPI**: `https://viacep.com.br/ws/{cep}/json`
