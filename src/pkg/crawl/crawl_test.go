package crawl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetItem(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		item, err := GetItem()
		assert.NoError(t, err)
		assert.NotEmpty(t, item)
	})
}
