package pokeapi

import (
	"encoding/json"
	"io"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, ok := c.cache.Get(url)
	if !ok {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		c.cache.Add(url, body)
	}

	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(body, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
