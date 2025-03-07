package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ExploreLocationArea(area string) ([]string, error) {
	fullURL := baseURL + "/location-area/" + area
	res, err := http.Get(fullURL)
	if err != nil {
		return []string{}, err
	}

	poke, ok := getExploreCacheCheck(c, fullURL)
	if ok {
		return poke, nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}

	c.Cache.Add(fullURL, body)

	var pokemon ExploreAreasResponse
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return []string{}, nil
	}

	var pokes []string
	for _, mon := range pokemon.Encounters {
		pokes = append(pokes, mon.Pokemon.Name)
	}

	return pokes, nil
}

func getExploreCacheCheck(c *Client, fullUrl string) ([]string, bool) {
	result, ok := c.Cache.Get(fullUrl)
	if ok {
		var pokemon ExploreAreasResponse
		err := json.Unmarshal(result, &pokemon)
		if err != nil {
			log.Printf("[CACHE] Miss to process: %s", fullUrl)
			return []string{}, false
		}

		var pokes []string
		for _, mon := range pokemon.Encounters {
			pokes = append(pokes, mon.Pokemon.Name)
		}

		log.Printf("[CACHE] Hit for URL: %s", fullUrl)

		return pokes, true
	}
	return []string{}, false
}
