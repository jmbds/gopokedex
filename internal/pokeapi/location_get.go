package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	var err error

	data, found := c.cache.Get(url)
	if !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Location{}, err
		}

		c.cache.Add(url, data)
	}

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	return locationResp, nil
}
