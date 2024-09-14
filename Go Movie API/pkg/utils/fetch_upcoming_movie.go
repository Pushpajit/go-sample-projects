package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"movie-api/pkg/structure"
	"net/http"
)

func GetUpcomingMovie(region string) structure.Response {
	var response structure.Response

	// endpoint to fetch the popular
	url := "https://api.themoviedb.org/3/movie/upcoming?language=en-US&page=1"
	if region != "" {
		url += fmt.Sprintf("&region=%v", region)
	}

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
