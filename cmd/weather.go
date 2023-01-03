package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weather"
)

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)

	apikey := os.Getenv("OWM_APIKEY")
	if apikey == "" {
		log.Fatal("please set the environment variable OWM_APIKEY")
	}

	if len(os.Args) < 2 {
		log.Fatalf("please supply city name")
	}
	city := os.Args[1]

	URL := weather.FormatURL(city, apikey)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected response status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	conditions, err := weather.ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.0f°C, %s\n", conditions.Temperature, conditions.Summary)
}
