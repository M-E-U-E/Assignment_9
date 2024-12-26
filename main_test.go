package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestSetupApplication(t *testing.T) {
	setupApplication()

	t.Run("Test Static File Headers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/static/test.css", nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, req)

		assert.Equal(t, "no-cache, no-store, must-revalidate", w.Header().Get("Cache-Control"))
		assert.Equal(t, "no-cache", w.Header().Get("Pragma"))
		assert.Equal(t, "0", w.Header().Get("Expires"))
	})

	t.Run("Test Template Function Directly", func(t *testing.T) {
		favFunc := func(input string) string {
			return fmt.Sprintf("Favorite ID: %s", input)
		}
		result := favFunc("123")
		assert.Equal(t, "Favorite ID: 123", result)
	})
}
