package discord

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDiscord_GetProperties(t *testing.T) {
	url := os.Getenv("TARGET_URL")
	d := NewDiscord(url)
	t.Run("ok", func(t *testing.T) {
		got := d.GetProperties()
		assert.NotEmpty(t, got)
	})
}
