package crawl

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var (
	crawl *Crawl
	url   string
)

func TestMain(m *testing.M) {
	println("before all...")
	url = os.Getenv("TARGET_URL")
	if len(url) == 0 {
		log.Fatal("failed to load env")
	}
	crawl, _ = NewCrawl()

	code := m.Run()
	println("after all...")

	os.Exit(code)
}

func TestGetProperties(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		items, err := crawl.ConstructDataBank.GetProperties(url)
		assert.NoError(t, err)
		assert.NotEmpty(t, items)

	})
}
