package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
	// "github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Set up basic configuration for testing
	web.BConfig.AppName = "cat_api"
	web.BConfig.RunMode = "test"
	web.BConfig.WebConfig.AutoRender = false
}

// func setupTestContext(req *http.Request) (*httptest.ResponseRecorder, *context.Context) {
// 	w := httptest.NewRecorder()
// 	ctx := &context.Context{
// 		Input:  &context.BeegoInput{Context: &context.Context{}},
// 		Output: &context.BeegoOutput{Context: &context.Context{}},
// 	}
// 	ctx.Reset(w, req)
// 	return w, ctx
// }

func TestBreedsController_Get(t *testing.T) {
	t.Run("Successful Response", func(t *testing.T) {
		// Mock HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "mock-api-key", r.Header.Get("xassert-api-key"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{
					"id": "abys",
					"name": "Abyssinian",
					"origin": "Egypt",
					"description": "Active and playful"
				}
			]`))
		}))
		defer server.Close()

		// Set up configuration
		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		// Create request and context
		req, _ := http.NewRequest("GET", "/breeds", nil)
		w, ctx := setupTestContext(req)

		// Execute controller
		controller := &BreedsController{}
		controller.Init(ctx, "BreedsController", "Get", nil)
		controller.Get()

		// Verify response body
		response := w.Body.String()
		assert.NotEmpty(t, response)

		// Verify template name is set
		assert.Equal(t, "breeds.tpl", controller.TplName)

		// Verify breeds data is set
		breeds, ok := controller.Data["Breeds"].([]Breed)
		assert.True(t, ok)
		assert.Len(t, breeds, 1)
		assert.Equal(t, "Abyssinian", breeds[0].Name)
	})

	t.Run("Missing API Key Configuration", func(t *testing.T) {
		// Clear API key configuration
		web.AppConfig.Set("api_key", "")

		req, _ := http.NewRequest("GET", "/breeds", nil)
		w, ctx := setupTestContext(req)

		controller := &BreedsController{}
		controller.Init(ctx, "BreedsController", "Get", nil)
		controller.Get()

		// Verify response body
		response := w.Body.String()
		assert.NotEmpty(t, response)

		var responseData map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &responseData)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to fetch API key from configuration", responseData["error"])
	})

	t.Run("Missing Base URL Configuration", func(t *testing.T) {
		// Set API key but clear base URL
		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", "")

		req, _ := http.NewRequest("GET", "/breeds", nil)
		w, ctx := setupTestContext(req)

		controller := &BreedsController{}
		controller.Init(ctx, "BreedsController", "Get", nil)
		controller.Get()

		// Verify response body
		response := w.Body.String()
		assert.NotEmpty(t, response)

		var responseData map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &responseData)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to fetch API base URL from configuration", responseData["error"])
	})

	t.Run("API Error Response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/breeds", nil)
		w, ctx := setupTestContext(req)

		controller := &BreedsController{}
		controller.Init(ctx, "BreedsController", "Get", nil)
		controller.Get()

		// Verify response body
		response := w.Body.String()
		assert.NotEmpty(t, response)

		var responseData map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &responseData)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to fetch breed data", responseData["error"])
	})

	t.Run("Invalid JSON Response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`invalid json`))
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/breeds", nil)
		w, ctx := setupTestContext(req)

		controller := &BreedsController{}
		controller.Init(ctx, "BreedsController", "Get", nil)
		controller.Get()

		// Verify response body
		response := w.Body.String()
		assert.NotEmpty(t, response)

		var responseData map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &responseData)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to fetch breed data", responseData["error"])
	})

	t.Run("Timeout Error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(6 * time.Second) // Longer than the 5-second timeout
		}))
		defer server.Close()
	
		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)
	
		req, _ := http.NewRequest("GET", "/breeds", nil)
		w, ctx := setupTestContext(req)
	
		controller := &BreedsController{}
		controller.Init(ctx, "BreedsController", "Get", nil)
		controller.Get()
	
		// Verify response body
		response := w.Body.String()
		assert.NotEmpty(t, response)
	
		var responseData map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &responseData)
		assert.NoError(t, err)
		assert.Equal(t, "Request timed out", responseData["error"])
	})
	
}

