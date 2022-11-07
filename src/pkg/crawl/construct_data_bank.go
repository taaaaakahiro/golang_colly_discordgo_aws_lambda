package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"golang-aws-lambda/src/pkg/domain/entity"
	"strings"
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
	c.OnHTML("#container01 > table > tbody", func(e *colly.HTMLElement) {
		items := make(map[int][]string, 0)
		idx := 0
		e.ForEach("#container01 > table > tbody > tr", func(i int, element1 *colly.HTMLElement) {

			if i != 0 {
				element1.ForEach("#container01 > table > tbody > tr > td", func(i int, element2 *colly.HTMLElement) {
					if i == 0 {
						href, _ := element2.DOM.Find("a").Attr("href")
						trim := strings.Split(element2.Text, "更新")
						project := strings.TrimSpace(trim[1])
						items[idx] = append(items[idx], project)
						items[idx] = append(items[idx], url+href[1:])

					} else if i == 1 {
						items[idx] = append(items[idx], element2.Text)
					} else if i == 2 {
						items[idx] = append(items[idx], element2.Text)
					}

				})
			}
			idx++
		})

		// limit 10
		for i := 1; i < 11; i++ {
			p := &entity.Property{
				Title:     items[i][0],
				DetailUrl: items[i][1],
				Address:   items[i][2],
				Square:    items[i][3],
			}
			ps = append(ps, p)
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
