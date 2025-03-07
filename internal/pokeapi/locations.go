package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(Url *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if Url != nil {
		url = *Url
	}

	if val, ok := c.Cache.Get(url); ok {
		location := LocationAreasResponse{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			log.Printf("[CACHE] Missed on key: %s\n", url)
			return LocationAreasResponse{}, err
		}

		log.Printf("[CACHE] Hit on key: %s\n", url)
		return location, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.Cache.Add(url, body)

	location := LocationAreasResponse{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return location, nil
}
