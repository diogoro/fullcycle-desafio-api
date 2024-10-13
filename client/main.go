package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerResponse struct {
	Bid string `json: "bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*300)
	defer cancel()
	sr, err := getServerResponse(ctx)
	if err != nil {
		log.Println(err)
	}
	err = saveFile(*sr)
	if err != nil {
		log.Println(err)
	}

}

func getServerResponse(ctx context.Context) (*ServerResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var sr ServerResponse
	err = json.NewDecoder(res.Body).Decode(&sr)
	if err != nil {
		return nil, err
	}
	return &sr, nil
}

func saveFile(sr ServerResponse) error {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(fmt.Sprintf("Dólar: %v", sr.Bid)))
	if err != nil {
		return err
	}
	return nil
}
