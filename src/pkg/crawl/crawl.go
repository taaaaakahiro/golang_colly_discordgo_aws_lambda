package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"golang-aws-lambda/src/pkg/domain/entity"
)

func GetItem(url string) (*entity.Item, error) {
	// Target URL
	var item entity.Item
	// Instantiate default collector
	c := colly.NewCollector()

	// Extract title element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
		item = entity.Item{
			Title: e.Text,
		}
	})

	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})

	// Start scraping on https://XXX
	c.Visit(url)

	return &item, nil

}
