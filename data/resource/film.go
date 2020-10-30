package resource

type Film struct {
	Title        string     `json:"title"`
	EpisodeId    int        `json:"episode_id"`
	OpeningCrawl string     `json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	ReleaseDate  string     `json:"release_date"`
	Characters   []Person   `json:"characters"`
	Planets      []Planet   `json:"planets"`
	Starships    []Starship `json:"starships"`
	Vehicles     []Vehicle  `json:"vehicles"`
	Species      []Species  `json:"species"`
	Metadata
}
