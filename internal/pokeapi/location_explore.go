package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(name string) (SingleLocation, error) {
	url := baseURL + "/location-area/" + name

	val, ok := c.cache.Get(url)
	if ok {
		locationsResp := SingleLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return SingleLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return SingleLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SingleLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return SingleLocation{}, err
	}

	locationsResp := SingleLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return SingleLocation{}, err
	}
	c.cache.Add(url, dat)
	return locationsResp, nil
}
