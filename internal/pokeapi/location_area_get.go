package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// GetLocationArea -
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName

	body, ok := c.cache.Get(url)
	if !ok {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}
		c.cache.Add(url, body)
	}

	locationArea := LocationArea{}
	err := json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to unmarshal location area: %w", err)
	}

	return locationArea, nil
}
