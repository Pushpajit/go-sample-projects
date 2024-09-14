package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"movie-api/pkg/structure"
	"net/http"
	"strings"
)

func GetSearchMovie(query string) structure.Response {
	var response structure.Response

	if query == "" {
		return response
	}

	// endpoint to fetch the popular
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%v&include_adult=false&language=en-US&page=1", strings.Join(strings.Split(query, " "), "%20"))

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI5NGQ0MTEyMmE5YTY2NTEyMTdkNjkyYjIxMDk4Y2ZmMyIsIm5iZiI6MTcyMjY3MjgyMS4yOTQ1MzcsInN1YiI6IjY2YWRkZTZjNjNjNjIwZTIxNmIwYWQ3MCIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.8B53wsO2AKQSUP5G9BpBiY-2rSrozEAU8cxLKp4xLfA")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("Error:", err)
	}

	return response
}
