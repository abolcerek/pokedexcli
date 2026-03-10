package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"math/rand/v2"
)

type results struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
type MapResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Prevous string `json:"previous"`
	Res []results `json:"results"`
}

type ExploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}


type CatchReponse struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  any  `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"past_abilities"`
	PastStats []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Stats []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		} `json:"stats"`
	} `json:"past_stats"`
	PastTypes []any `json:"past_types"`
	Species   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       string `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationIx struct {
				ScarletViolet struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"scarlet-violet"`
			} `json:"generation-ix"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       string `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  string `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      string `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				BrilliantDiamondShiningPearl struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"brilliant-diamond-shining-pearl"`
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func cleanInput(text string)[]string {
	if len(text) == 0 {
		return []string{}
	}
	lowerstring := strings.ToLower(text)
	words := strings.Fields(lowerstring)
	return words
}

func commandExit(config *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Displays 20 location areas\nmapb: Displays previous 20 location areas\nexplore: Lists all pokemon in location\ncatch: Catches a pokemon\ninspect: Inspects a pokemon\npokedex: Displays pokedex\n")
	return nil
}
func commandMap(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != "" {
		url = config.Next
	}
	value, ok := config.cache.Get(url)
	if ok == true {
		res := MapResponse{}
		err := json.Unmarshal(value, &res)
		if err != nil {
			return err
		}
		for _, location := range res.Res {
			fmt.Printf("%s\n", location.Name)
		}
		config.Next = res.Next
		config.Previous = res.Prevous
		return nil
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return err
		}
		config.cache.Add(url, body)
		res := MapResponse{}
		err = json.Unmarshal(body, &res)
		if err != nil {
			return err
		}
		for _, location := range res.Res {
			fmt.Printf("%s\n", location.Name)
		}
		config.Next = res.Next
		config.Previous = res.Prevous
		return nil
	}
}
func commandMapb(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Previous != "" {
		url = config.Previous
	}
	value, ok := config.cache.Get(url)
	if ok == true {
		res := MapResponse{}
		err := json.Unmarshal(value, &res)
		if err != nil {
			return err
		}
		for _, location := range res.Res {
			fmt.Printf("%s\n", location.Name)
		}
		config.Next = res.Next
		config.Previous = res.Prevous
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	config.cache.Add(url, body)
	res := MapResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	if config.Previous == res.Prevous {
		fmt.Print("you're on the first page\n")
		config.Previous = ""
		config.Next = ""
		return nil
	}
	for _, location := range res.Res {
		fmt.Printf("%s\n", location.Name)
	}
	config.Next = res.Next
	config.Previous = res.Prevous
	return nil
}

func commandExplore(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/" + config.Location
	value, ok := config.cache.Get(url)
	if ok == true {
		res := ExploreResponse{}
		err := json.Unmarshal(value, &res)
		if err != nil {
			return err
		}
		for _, encounter := range res.PokemonEncounters {
			fmt.Printf("%s\n", encounter.Pokemon.Name)
		}
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	config.cache.Add(url, body)
	res := ExploreResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	fmt.Print("Exploring pastoria-city-area...\n")
	fmt.Print("Found Pokemon:\n")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *config) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + config.Pokemon
	value, ok := config.cache.Get(url)
	if ok == true {
		res := CatchReponse{}
		err := json.Unmarshal(value, &res)
		if err != nil {
			return err
		}
		fmt.Printf("Throwing a Pokeball at %s...\n", config.Pokemon)
		capture_rate := float64(res.BaseExperience) / 250
		if rand.Float64() > capture_rate {
			fmt.Printf("%s was caught!\n", config.Pokemon)
			var PokemonTypes []string
			for i:=0; i<len(res.Types); i++ {
				PokemonTypes = append(PokemonTypes, res.Types[i].Type.Name)
			}
			config.Pokedex[config.Pokemon] = Pokemon{
				name: config.Pokemon,
				height: res.Height,
				weight: res.Weight,
				Stats: PokemonStats{
					hp: res.Stats[0].BaseStat,
					attack: res.Stats[1].BaseStat,
					defense: res.Stats[2].BaseStat,
					special_attack: res.Stats[3].BaseStat,
					special_defense: res.Stats[4].BaseStat,
					speed: res.Stats[5].BaseStat,
				},
				Types: PokemonTypes,
			}
		} else {
			fmt.Printf("%s escaped!\n", config.Pokemon)
		}
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	config.cache.Add(url, body)
	res := CatchReponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", config.Pokemon)
	capture_rate := float64(res.BaseExperience) / 250
	if rand.Float64() > capture_rate {
		fmt.Printf("%s was caught!\n", config.Pokemon)
		var PokemonTypes []string
		for i:=0; i<len(res.Types); i++ {
			PokemonTypes = append(PokemonTypes, res.Types[i].Type.Name)
		}
		config.Pokedex[config.Pokemon] = Pokemon{
			name: config.Pokemon,
			height: res.Height,
			weight: res.Weight,
			Stats: PokemonStats{
				hp: res.Stats[0].BaseStat,
				attack: res.Stats[1].BaseStat,
				defense: res.Stats[2].BaseStat,
				special_attack: res.Stats[3].BaseStat,
				special_defense: res.Stats[4].BaseStat,
				speed: res.Stats[5].BaseStat,
			},
			Types: PokemonTypes,
		}
	} else {
		fmt.Printf("%s escaped!\n", config.Pokemon)
	}
	return nil
}

func commandInspect(config *config) error {
	fmt.Printf("Name: %s\n", config.Pokedex[config.Pokemon].name)
	fmt.Printf("Height: %d\n", config.Pokedex[config.Pokemon].height)
	fmt.Printf("Weight: %d\n", config.Pokedex[config.Pokemon].weight)
	fmt.Print("Stats:\n")
	fmt.Printf(" -hp: %d\n", config.Pokedex[config.Pokemon].Stats.hp)
	fmt.Printf(" -attack: %d\n", config.Pokedex[config.Pokemon].Stats.attack)
	fmt.Printf(" -defense: %d\n", config.Pokedex[config.Pokemon].Stats.defense)
	fmt.Printf(" -special-attack: %d\n", config.Pokedex[config.Pokemon].Stats.special_attack)
	fmt.Printf(" -special-defense: %d\n", config.Pokedex[config.Pokemon].Stats.special_defense)
	fmt.Printf(" -speed: %d\n", config.Pokedex[config.Pokemon].Stats.speed)
	fmt.Print("Types:\n")
	for i := 0; i<len(config.Pokedex[config.Pokemon].Types); i++ {
		fmt.Printf(" - %s\n", config.Pokedex[config.Pokemon].Types[i])
	}
	return nil
}

func commandPokedex(config *config) error {
	fmt.Print("Your Pokedex:\n")
	for index, _ := range config.Pokedex{
		fmt.Printf(" - %s\n", config.Pokedex[index].name)
	}
	return nil
}

