package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type AnimeNews struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CreatedAt   string   `json:"creatdat"`
	CreatedBy   string   `json:"createdby"`
	Tags        []string `json:"tags"`
}

var anime []AnimeNews

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("myanimelist.net"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnHTML("div[class='news-unit clearfix rect']", func(h *colly.HTMLElement) {
		// html, _ := h.DOM.Html()
		title := h.ChildText("div[class='news-unit-right'] p.title a")
		desc := h.ChildText("div[class='news-unit-right'] div.text")
		createdAt := h.ChildText("div[class='news-unit-right'] div.information p[class='info di-ib']")

		var (
			craetedBy []string
			tags      []string
		)

		h.ForEach("div[class='news-unit-right'] div.information a", func(i int, h *colly.HTMLElement) {
			craetedBy = append(craetedBy, h.Text)
		})

		h.ForEach("div[class='news-unit-right'] div.information p[class='di-ib tags']", func(i int, h *colly.HTMLElement) {
			tags = append(tags, h.Text)
		})

		// removing redundant string and info
		if len(craetedBy) > 1 {
			createdAt = strings.Replace(createdAt, craetedBy[0], "", 1)
			createdAt = strings.Replace(createdAt, craetedBy[1], "", 1)
			createdAt = strings.Replace(createdAt, "by ", "", 1)
			createdAt = strings.Trim(createdAt, " ")
		}

		// Added to the animenews slice
		anime = append(anime, AnimeNews{
			Title:       title,
			Description: desc,
			CreatedBy:   craetedBy[0],
			CreatedAt:   createdAt,
			Tags:        tags,
		})

		// fmt.Printf("Title: %v\nDesc: %v\n CreatedBy: %v\n CreatedAt: %v\nTags: %+v\n\n", title, desc, craetedBy[0], createdAt, tags)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Scrapping the site: %v\n", r.URL)
	})

	c.Visit("https://myanimelist.net/news")

	bytes, _ := json.MarshalIndent(anime, "", " ")
	if err := os.WriteFile("animenews.json", bytes, 0644); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Total %v news are fetched\n", len(anime))
	fmt.Println("Anime news JSON file has been generated!")

}
