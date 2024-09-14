package main

import (
	"fmt"
	"movie-api/pkg/structure"
	"movie-api/pkg/utils"
)

func main() {

	response := utils.GetPopularMovies("")

	for _, movie := range response.Results {
		fmt.Println("ID:", movie.Id)
		fmt.Println("Title:", movie.Title)
		fmt.Println("Overview:", movie.Overview)
		fmt.Println("Release Date:", movie.Date)
		fmt.Println("Rating:", movie.Rating)
		fmt.Println("Poster Path:", movie.Poster)
		fmt.Println("Backdrop Path:", movie.Backdrop)
		fmt.Printf("Genres: ")
		for _, v := range movie.Genres {
			fmt.Printf("%v ", structure.MovieGenre[v])
		}
		fmt.Println()
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Total Item Fetched:", len(response.Results))

}
