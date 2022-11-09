package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"golang-aws-lambda/src/pkg/domain/entity"
	"strconv"
	"strings"
	"time"
)

type ConstructDataBank struct {
	colly *colly.Collector
}

func NewConstructDataBank(c *colly.Collector) *ConstructDataBank {
	return &ConstructDataBank{
		colly: c,
	}
}

func (cdb *ConstructDataBank) GetProperties(url string) ([]*entity.Property, error) {
	ps := make([]*entity.Property, 0)
	c := cdb.colly

	// start crawl
	var (
		items   []string
		updates []string
		urls    []string
		address []string
		square  []string
	)
	c.OnHTML("#container01 > table > tbody", func(e *colly.HTMLElement) {
		idx := 0
		e.ForEach("#container01 > table > tbody > tr", func(i int, element1 *colly.HTMLElement) {
			if i != 0 {
				element1.ForEach("#container01 > table > tbody > tr > td", func(i2 int, element2 *colly.HTMLElement) {
					if i2 == 0 {
						href, _ := element2.DOM.Find("a").Attr("href")
						p := strings.Split(element2.Text, "更新")
						project := strings.TrimSpace(p[1])
						items = append(items, project)
						urls = append(urls, url+href[1:])
						updates = append(updates, p[0])

					}
					if i2 == 1 {
						address = append(address, element2.Text)
					}
					if i2 == 2 {
						square = append(square, element2.Text)
					}
				})
			}
			idx++
		})

		// limit 10
		for i := 0; i < 10; i++ {
			if isToday(updates[i]) {
				p := &entity.Property{
					Title:     items[i],
					Updated:   updates[i],
					DetailUrl: urls[i],
					Address:   address[i],
					Square:    square[i],
				}
				ps = append(ps, p)
			}
		}

	})

	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})

	// Start scraping on https://XXX
	c.Visit(url)

	return ps, nil

}

func isToday(strDate string) bool {
	today := time.Now()
	mon := int(today.Month())
	day := today.Day()
	trim := strings.Split(strDate, "/")
	getMon, err := strconv.Atoi(trim[0])
	if err != nil {
		return false
	}
	getDay, err := strconv.Atoi(trim[1])
	if err != nil {
		return false
	}
	if getMon == mon && getDay == day {
		return true
	}
	return false

}
