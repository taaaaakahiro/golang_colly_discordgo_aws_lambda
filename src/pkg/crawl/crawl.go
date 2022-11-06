package crawl

import "github.com/gocolly/colly"

type Crawl struct {
	ConstructDataBank *ConstructDataBank
}

func NewCrawl() (*Crawl, error) {
	// init colly
	c := colly.NewCollector()
	return &Crawl{
		ConstructDataBank: NewConstructDataBank(c),
	}, nil
}
