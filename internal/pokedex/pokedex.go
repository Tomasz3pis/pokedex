package pokedex

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"ability,omitempty"`
		IsHidden bool `json:"is_hidden,omitempty"`
		Slot     int  `json:"slot,omitempty"`
	} `json:"abilities,omitempty"`
	BaseExperience int `json:"base_experience,omitempty"`
	Cries          struct {
		Latest string `json:"latest,omitempty"`
		Legacy string `json:"legacy,omitempty"`
	} `json:"cries,omitempty"`
	Forms []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"forms,omitempty"`
	GameIndices []struct {
		GameIndex int `json:"game_index,omitempty"`
		Version   struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"version,omitempty"`
	} `json:"game_indices,omitempty"`
	Height    int `json:"height,omitempty"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"item,omitempty"`
		VersionDetails []struct {
			Rarity  int `json:"rarity,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"held_items,omitempty"`
	ID                     int    `json:"id,omitempty"`
	IsDefault              bool   `json:"is_default,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"move,omitempty"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at,omitempty"`
			MoveLearnMethod struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"move_learn_method,omitempty"`
			VersionGroup struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version_group,omitempty"`
		} `json:"version_group_details,omitempty"`
	} `json:"moves,omitempty"`
	Name          string `json:"name,omitempty"`
	Order         int    `json:"order,omitempty"`
	PastAbilities []any  `json:"past_abilities,omitempty"`
	PastTypes     []any  `json:"past_types,omitempty"`
	Species       struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"species,omitempty"`
	Sprites struct {
		BackDefault      string `json:"back_default,omitempty"`
		BackFemale       string `json:"back_female,omitempty"`
		BackShiny        string `json:"back_shiny,omitempty"`
		BackShinyFemale  string `json:"back_shiny_female,omitempty"`
		FrontDefault     string `json:"front_default,omitempty"`
		FrontFemale      string `json:"front_female,omitempty"`
		FrontShiny       string `json:"front_shiny,omitempty"`
		FrontShinyFemale string `json:"front_shiny_female,omitempty"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontFemale  any    `json:"front_female,omitempty"`
			} `json:"dream_world,omitempty"`
			Home struct {
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      string `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale string `json:"front_shiny_female,omitempty"`
			} `json:"home,omitempty"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
			} `json:"official-artwork,omitempty"`
			Showdown struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       string `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      string `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale string `json:"front_shiny_female,omitempty"`
			} `json:"showdown,omitempty"`
		} `json:"other,omitempty"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackGray         string `json:"back_gray,omitempty"`
					BackTransparent  string `json:"back_transparent,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontGray        string `json:"front_gray,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"red-blue,omitempty"`
				Yellow struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackGray         string `json:"back_gray,omitempty"`
					BackTransparent  string `json:"back_transparent,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontGray        string `json:"front_gray,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"yellow,omitempty"`
			} `json:"generation-i,omitempty"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default,omitempty"`
					BackShiny             string `json:"back_shiny,omitempty"`
					BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
					BackTransparent       string `json:"back_transparent,omitempty"`
					FrontDefault          string `json:"front_default,omitempty"`
					FrontShiny            string `json:"front_shiny,omitempty"`
					FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
					FrontTransparent      string `json:"front_transparent,omitempty"`
				} `json:"crystal,omitempty"`
				Gold struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"gold,omitempty"`
				Silver struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"silver,omitempty"`
			} `json:"generation-ii,omitempty"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"emerald,omitempty"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"firered-leafgreen,omitempty"`
				RubySapphire struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"ruby-sapphire,omitempty"`
			} `json:"generation-iii,omitempty"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"diamond-pearl,omitempty"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"heartgold-soulsilver,omitempty"`
				Platinum struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"platinum,omitempty"`
			} `json:"generation-iv,omitempty"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default,omitempty"`
						BackFemale       string `json:"back_female,omitempty"`
						BackShiny        string `json:"back_shiny,omitempty"`
						BackShinyFemale  string `json:"back_shiny_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						FrontFemale      string `json:"front_female,omitempty"`
						FrontShiny       string `json:"front_shiny,omitempty"`
						FrontShinyFemale string `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"black-white,omitempty"`
			} `json:"generation-v,omitempty"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"omegaruby-alphasapphire,omitempty"`
				XY struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"x-y,omitempty"`
			} `json:"generation-vi,omitempty"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  any    `json:"front_female,omitempty"`
				} `json:"icons,omitempty"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"ultra-sun-ultra-moon,omitempty"`
			} `json:"generation-vii,omitempty"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  string `json:"front_female,omitempty"`
				} `json:"icons,omitempty"`
			} `json:"generation-viii,omitempty"`
		} `json:"versions,omitempty"`
	} `json:"sprites,omitempty"`
	Stats []struct {
		BaseStat int `json:"base_stat,omitempty"`
		Effort   int `json:"effort,omitempty"`
		Stat     struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot int `json:"slot,omitempty"`
		Type struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"type,omitempty"`
	} `json:"types,omitempty"`
	Weight int `json:"weight,omitempty"`
}
type Stats struct {
	Name  string
	Value int
}
type Types struct {
	Name string
}
type InspectInfo struct {
	Name   string
	Height int
	Weight int
	Stats  []Stats
	Types  []Types
}
type Pokedex struct {
	entry map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		make(map[string]Pokemon),
	}
}

func (pokedex *Pokedex) AddEntry(key string, pokemon Pokemon) {
	pokedex.entry[key] = pokemon
}

func (pokedex *Pokedex) GetEntry(key string) (InspectInfo, bool) {
	if v, exists := pokedex.entry[key]; exists {

		st := []Stats{}
		ty := []Types{}
		for _, v := range v.Stats {
			st = append(st, Stats{Name: v.Stat.Name, Value: v.BaseStat})
		}
		for _, v := range v.Types {
			ty = append(ty, Types{Name: v.Type.Name})
		}
		return InspectInfo{
			Name:   v.Name,
			Height: v.Height,
			Weight: v.Weight,
			Stats:  st,
			Types:  ty,
		}, true

	} else {
		return InspectInfo{}, false
	}
}

func (pokedex *Pokedex) ListEntries() []string {
	entries := []string{}
	for _, v := range pokedex.entry {
		entries = append(entries, v.Name)
	}
	return entries
}
