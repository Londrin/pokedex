package pokeapi

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type ExploreAreasResponse struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Base_experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Is_default      bool   `json:"is_default"`
	Order           int    `json:"order"`
	Weight          int    `json:"weight"`
	Abilities       []struct {
		Is_hidden bool `json:"is_hidden"`
		Slot      int  `json:"slot"`
		Ability   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`
	Game_indices []struct {
		Game_index int `json:"game_index"`
		Version    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Held_items []struct {
		Item struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"item"`
		Version_details []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
		Version_group_details []struct {
			Level_learned_at int `json:"level_learned_at"`
			Version_group    struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version_group"`
			Move_learn_method struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		Base_stat int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
