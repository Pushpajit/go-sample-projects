package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var (
	BaseURL = "https://codeforces.com/problemset"
)

type Problem struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Topics     []string `json:"topics"`
	SolvedBy   int      `json:"solved-by"`
	Difficulty int      `json:"difficulty"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("codeforces.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// set a global storage for answer.
	var problems []Problem

	c.OnHTML("table.problems tbody tr", func(h *colly.HTMLElement) {
		var diff int
		var elements []string
		var problem Problem

		id := h.ChildText("td.id a")
		problem.Id = id

		if h.ChildText("td span") != "" {
			diff, _ = strconv.Atoi(h.ChildText("td span")) // strin to interger conversion
		}

		// This part is tricky because just look at the HTML
		h.ForEach("td div a", func(i int, h *colly.HTMLElement) {
			elements = append(elements, h.Text)
		})

		// check the corner case, if there is no element found
		if len(elements) > 0 {
			problem.Name = strings.Trim(strings.Trim(elements[0], "\n"), " ")
			problem.Topics = elements[1:]
		}

		problem.Difficulty = diff // Fetch the difficulty lvl
		
		solvedby := h.ChildText("td a[title='Participants solved the problem']")
		if len(solvedby) > 0 {
			s, _ := strconv.Atoi(solvedby[1:])
			problem.SolvedBy = s
		} else {
			s, _ := strconv.Atoi(solvedby)
			problem.SolvedBy = s
		}

		problems = append(problems, problem)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visit: %v\n", r.URL)
	})

	c.Visit(BaseURL) // Start the process

	// Saving the data into a JSON file
	res, _ := json.MarshalIndent(problems, "", " ")
	fmt.Println(string(res))
	os.WriteFile("codeforces-problem.json", res, 0644)

	
}
