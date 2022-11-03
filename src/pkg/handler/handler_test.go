package handler

import (
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		hook := os.Getenv("HOOK_ESC_KEY")
		Handler(hook)

	})
}
