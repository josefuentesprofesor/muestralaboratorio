package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PriceResponse struct {
	Time       TimeData `json:"time"`
	Disclaimer string   `json:"disclaimer"`
	BPI        BPIData  `json:"bpi"`
}
type TimeData struct {
	Updated    string `json:"updated"`
	UpdatedISO string `json:"updatedISO"`
	UpdatedUK  string `json:"updateduk"`
}

type BPIData struct {
	USD Currency `json:"USD"`
	BTC Currency `json:"BTC"`
}

type Currency struct {
	Code        string  `json:"code"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

func GetBitcoinPrice() string {
	url := "https://api.coindesk.com/v1/bpi/currentprice/BTC.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud HTTP:", err)
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta HTTP:", err)
		return ""
	}
	var priceData PriceResponse
	err = json.Unmarshal(body, &priceData)
	if err != nil {
		fmt.Println("Error al decodificar el JSON:", err)
		return ""
	}
	btcPrice := priceData.BPI.USD.RateFloat
	resp.Body.Close()
	result := fmt.Sprint("USD ", btcPrice)
	return result
}
