package movie

type Movies struct {
	Dates         Date    `json:"dates"`
	Page          int64   `json:"page"`
	Results       []Movie `json:"results"`
	Total_pages   int64   `json:"total_pages"`
	Total_Results int64   `json:"total_results"`
}

type Movie struct {
	Adult        string `json:"adult"`
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	Release_date string `json:"release_date"`
	Video        bool   `json:"video"`
	Vote_average string `json:"vote_average"`
}

type Date struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}
