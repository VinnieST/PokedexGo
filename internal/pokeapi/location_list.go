package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (*LocationListResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.Cache.Get(url); ok {
		locationsResp := &LocationListResponse{}
		if err := json.Unmarshal(data, locationsResp); err == nil {
			return locationsResp, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	locationsResp := &LocationListResponse{}
	err = json.Unmarshal(dat, locationsResp)
	if err != nil {
		return nil, err
	}

	return locationsResp, nil
}
