package data

type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	AverageLifespan string   `json:"average_lifespan"`
	HairColors      string   `json:"hair_colors"`
	SkinColors      string   `json:"skin_colors"`
	EyeColors       string   `json:"eye_colors"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []Person `json:"people"`
	Films           []Film   `json:"films"`
	Metadata
}
