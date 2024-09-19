package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

type Weather struct {
	Description string `json:"description"`
	Days        []Day  `json:"days"`
}

type Day struct {
	Datetime          string  `json:"datetime"`
	TempMax           float64 `json:"tempmax"`
	TempMin           float64 `json:"tempmin"`
	Conditions        string  `json:"conditions"`
	PrecipitationProb float64 `json:"precipprob"`
}

func main() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

	url := os.Getenv("API_URL")
	if url == "" {
		fmt.Println("Error: API_URL environment variable is not set")
		return
	}
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather API is not available")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	if len(weather.Days) > 0 {
		today := weather.Days[0]
		fmt.Println("Date:", today.Datetime)
		fmt.Println("Max Temperature:", today.TempMax)
		fmt.Println("Min Temperature:", today.TempMin)
		fmt.Println("Conditions:", today.Conditions)
		fmt.Println("Precipitation Probability:", today.PrecipitationProb)
	}
}
