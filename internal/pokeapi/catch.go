package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetCatchPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.Cache.Get(url); ok {
		poke := Pokemon{}
		err := json.Unmarshal(val, &poke)
		if err != nil {
			return Pokemon{}, err
		}

		return poke, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	poke := Pokemon{}
	err = json.Unmarshal(body, &poke)
	if err != nil {
		return Pokemon{}, err
	}

	c.Cache.Add(url, body)

	return poke, nil
}
