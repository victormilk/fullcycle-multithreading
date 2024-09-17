package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	BRASIL_API      = "https://brasilapi.com.br/api/cep/v1"
	VIACEP_API      = "https://viacep.com.br/ws"
	REQUEST_TIMEOUT = time.Second * 1
	DEFAULT_CEP     = "01153000"
)

type Winner struct {
	Name    string
	Payload any
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	defer cancel()

	channel := make(chan *Winner, 1)

	var cep string
	if len(os.Args) > 1 {
		cep = os.Args[1]
	} else {
		cep = DEFAULT_CEP
	}

	go func() {
		payload, err := getBrasilAPI(ctx, cep)
		if err != nil {
			fmt.Println(err)
			return
		}
		channel <- &Winner{
			Name:    "BrasilAPI",
			Payload: payload,
		}
	}()

	go func() {
		payload, err := getViaCepAPI(ctx, cep)
		if err != nil {
			fmt.Println(err)
			return
		}

		channel <- &Winner{
			Name:    "ViaCepAPI",
			Payload: payload,
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Request timeout exceeded")
		return
	case winner := <-channel:
		fmt.Printf("Winner: %s\n", winner.Name)
		fmt.Printf("Payload: %v\n", winner.Payload)
		return
	}

}

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func getBrasilAPI(ctx context.Context, cep string) (string, error) {
	path := fmt.Sprintf("%s/%s", BRASIL_API, cep)
	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getViaCepAPI(ctx context.Context, cep string) (string, error) {
	path := fmt.Sprintf("%s/%s/json", VIACEP_API, cep)
	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
