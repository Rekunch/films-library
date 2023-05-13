package kinopoisk_dev

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Rekunch/films-library/internal/model"
)

func Random() string {
	apiKey := os.Getenv("API_KEY")

	req, _ := http.NewRequest("GET", RandomMovie, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var movie model.Movie

	err := json.Unmarshal(body, &movie)
	if err != nil {
		return ""
	}

	var genreNames []string
	for _, genre := range movie.Genres {
		genreNames = append(genreNames, genre.Name)
	}
	joinedGenres := strings.Join(genreNames, ", ")

	var countryNames []string
	for _, country := range movie.Countries {
		countryNames = append(countryNames, country.Name)
	}
	joinedCountries := strings.Join(countryNames, ", ")

	movieInfo := fmt.Sprintf("Название фильма: %s\nЖанр: %s\nСтрана: %s\nДлительность: %d минут\nОписание:%s\nРейтинг Imdb: %.2f\nРейтинг Кинопоиска: %.2f\nГод выпуска: %d\n%s", movie.Name, joinedGenres, joinedCountries, movie.Length, movie.Description, movie.Rating.Imdb, movie.Rating.Kp, movie.Year, movie.Poster.PreviewUrl)

	return movieInfo
}

func FindMovieByName(name string) []string {
	apiKey := os.Getenv("API_KEY")

	//url := fmt.Sprintf("%s%s", MovieByName, name)
	nameMovie := url.QueryEscape(name)

	req, _ := http.NewRequest("GET", MovieByName+nameMovie, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var movie model.Response

	err := json.Unmarshal(body, &movie)
	if err != nil {
		return nil
	}

	var genreNames []string
	var countryNames []string
	var movieInfoList []string

	for i := 0; i < len(movie.Docs); i++ {
		for _, genre := range movie.Docs[i].Genres {
			genreNames = append(genreNames, genre.Name)
		}
		joinedGenres := strings.Join(genreNames, ", ")
		for _, country := range movie.Docs[i].Countries {
			countryNames = append(countryNames, country.Name)
		}
		joinedCountries := strings.Join(countryNames, ", ")
		movieInfo := fmt.Sprintf("Название фильма: %s\nЖанр: %s\nСтрана: %s\nДлительность: %d минут\nОписание:%s\nРейтинг Imdb: %.2f\nРейтинг Кинопоиска: %.2f\nГод выпуска: %d\n%s\n\n", movie.Docs[i].Name, joinedGenres, joinedCountries, movie.Docs[i].Length, movie.Docs[i].Description, movie.Docs[i].Rating.Imdb, movie.Docs[i].Rating.Kp, movie.Docs[i].Year, movie.Docs[i].Poster.PreviewUrl)
		movieInfoList = append(movieInfoList, movieInfo)

		// clear genres and countries for the next iteration
		genreNames = nil
		countryNames = nil
	}

	return movieInfoList
}
