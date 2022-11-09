package crawl

import (
	"github.com/stretchr/testify/assert"
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
