package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getLocationAreas() (LocationArea, error) {
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LocationArea{}, fmt.Errorf("unexpected response status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	return locationArea, nil
}

func commandMap(config *LocationArea) error {

	for _, result := range config.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
