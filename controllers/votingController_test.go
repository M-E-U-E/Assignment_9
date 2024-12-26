package controllers

import (
	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestVotingController_Get(t *testing.T) {
	t.Run("Successful Response", func(t *testing.T) {
		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", "http://mock-api.com")

		// Mock HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "mock-api-key", r.Header.Get("x-api-key"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":"1","url":"http://example.com/cat.jpg"}]`))
		}))
		defer server.Close()

		web.AppConfig.Set("api_base_url", server.URL)

		// Create mock Beego context
		req, _ := http.NewRequest("GET", "/voting", nil)
		w := httptest.NewRecorder()
		ctx := &context.Context{
			Input:  &context.BeegoInput{Context: &context.Context{}},
			Output: &context.BeegoOutput{Context: &context.Context{}},
		}
		ctx.Reset(w, req)

		// Initialize and execute controller
		controller := &VotingController{}
		controller.Init(ctx, "VotingController", "Get", nil)
		controller.Get()

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "http://example.com/cat.jpg")
	})

	t.Run("API Error", func(t *testing.T) {
		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", "http://mock-api.com")

		// Mock HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/voting", nil)
		w := httptest.NewRecorder()
		ctx := &context.Context{
			Input:  &context.BeegoInput{Context: &context.Context{}},
			Output: &context.BeegoOutput{Context: &context.Context{}},
		}
		ctx.Reset(w, req)

		controller := &VotingController{}
		controller.Init(ctx, "VotingController", "Get", nil)
		controller.Get()

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to fetch data")
	})

	t.Run("Timeout Error", func(t *testing.T) {
		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", "http://mock-api.com")

		// Mock HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(6 * time.Second)
		}))
		defer server.Close()

		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/voting", nil)
		w := httptest.NewRecorder()
		ctx := &context.Context{
			Input:  &context.BeegoInput{Context: &context.Context{}},
			Output: &context.BeegoOutput{Context: &context.Context{}},
		}
		ctx.Reset(w, req)

		controller := &VotingController{}
		controller.Init(ctx, "VotingController", "Get", nil)
		controller.Get()

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)
		assert.Contains(t, w.Body.String(), "Request timed out")
	})
}
