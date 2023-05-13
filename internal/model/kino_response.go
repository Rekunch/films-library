package model

type Response struct {
	Docs  []Movie `json:"docs"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
	Page  int     `json:"page"`
	Pages int     `json:"pages"`
}

type Movie struct {
	Name             string    `json:"name"`
	Genres           []Genre   `json:"genres"`
	Countries        []Country `json:"countries"`
	Length           int64     `json:"movieLength"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"shortDescription"`
	Rating           struct {
		Kp   float64 `json:"kp"`
		Imdb float64 `json:"imdb"`
	} `json:"rating"`
	Year   int64 `json:"year"`
	Poster struct {
		PreviewUrl string `json:"previewUrl"`
	} `json:"poster"`
}

type Genre struct {
	Name string `json:"name"`
}

type Country struct {
	Name string `json:"name"`
}
