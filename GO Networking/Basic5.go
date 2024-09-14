package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type todos struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}

func fetch(url string, method string) []todos {
	var content []todos

	if method == "GET" {
		resp, err := http.Get(url)
		checkNil(err)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			panic(fmt.Sprintf("failed to fetch data: status code %d", resp.StatusCode))
		}

		bytecode, _ := io.ReadAll(resp.Body)
		checkNil(err)

		json.Unmarshal(bytecode, &content)

	}

	return content
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"
	res := fetch(url, "GET")

	for _, val := range res {
		fmt.Printf("ID: %v\nTitle: %v\n\n", val.Id, val.Title)
	}
}
