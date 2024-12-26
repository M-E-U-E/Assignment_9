// package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/beego/beego/v2/server/web"
// 	// "github.com/beego/beego/v2/server/web/context"
// 	"github.com/stretchr/testify/assert"
// )

// // func setupTestContext(req *http.Request) (*httptest.ResponseRecorder, *context.Context) {
// // 	w := httptest.NewRecorder()
// // 	ctx := &context.Context{
// // 		Input:  &context.BeegoInput{Context: &context.Context{}},
// // 		Output: &context.BeegoOutput{Context: &context.Context{}},
// // 	}
// // 	ctx.Reset(w, req)
// // 	return w, ctx
// // }

// func TestFavsController_Get(t *testing.T) {
// 	t.Run("Successful Response", func(t *testing.T) {
// 		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte(`[{"id":1,"image":{"id":"cat1","url":"http://example.com/cat1.jpg"}}]`))
// 		}))
// 		defer server.Close()

// 		web.AppConfig.Set("api_key", "mock-api-key")
// 		web.AppConfig.Set("api_base_url", server.URL)

// 		req, _ := http.NewRequest("GET", "/favs", nil)
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Get", nil)
// 		controller.Get()

// 		assert.NotEmpty(t, w.Body.String())
// 		assert.Contains(t, w.Body.String(), "http://example.com/cat1.jpg")
// 	})

// 	t.Run("Missing API Key Configuration", func(t *testing.T) {
// 		web.AppConfig.Set("api_key", "")

// 		req, _ := http.NewRequest("GET", "/favs", nil)
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Get", nil)
// 		controller.Get()

// 		assert.NotEmpty(t, w.Body.String())
// 		assert.Contains(t, w.Body.String(), "Failed to fetch API key from configuration")
// 	})

// 	t.Run("Error in Fetching Favorites", func(t *testing.T) {
// 		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		}))
// 		defer server.Close()

// 		web.AppConfig.Set("api_key", "mock-api-key")
// 		web.AppConfig.Set("api_base_url", server.URL)

// 		req, _ := http.NewRequest("GET", "/favs", nil)
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Get", nil)
// 		controller.Get()

// 		assert.NotEmpty(t, w.Body.String())
// 		assert.Contains(t, w.Body.String(), "Error fetching favorites")
// 	})

// 	t.Run("Timeout Scenario", func(t *testing.T) {
// 		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			time.Sleep(6 * time.Second)
// 		}))
// 		defer server.Close()

// 		web.AppConfig.Set("api_key", "mock-api-key")
// 		web.AppConfig.Set("api_base_url", server.URL)

// 		req, _ := http.NewRequest("GET", "/favs", nil)
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Get", nil)
// 		controller.Get()

// 		assert.NotEmpty(t, w.Body.String())
// 		assert.Contains(t, w.Body.String(), "Request timed out")
// 	})
// }

// func TestFavsController_Post(t *testing.T) {
// 	t.Run("Successful Post", func(t *testing.T) {
// 		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte(`{"message":"SUCCESS"}`))
// 		}))
// 		defer server.Close()

// 		web.AppConfig.Set("api_key", "mock-api-key")
// 		web.AppConfig.Set("api_base_url", server.URL)

// 		payload := map[string]string{"image_id": "cat1"}
// 		payloadJSON, _ := json.Marshal(payload)
// 		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer(payloadJSON))
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Post", nil)
// 		controller.Post()

// 		var response map[string]string
// 		err := json.Unmarshal(w.Body.Bytes(), &response)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "SUCCESS", response["message"])
// 	})

// 	t.Run("Missing API Key Configuration", func(t *testing.T) {
// 		web.AppConfig.Set("api_key", "")

// 		payload := map[string]string{"image_id": "cat1"}
// 		payloadJSON, _ := json.Marshal(payload)
// 		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer(payloadJSON))
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Post", nil)
// 		controller.Post()

// 		assert.Contains(t, w.Body.String(), "Failed to fetch API key from configuration")
// 	})

// 	t.Run("Invalid Request Body", func(t *testing.T) {
// 		web.AppConfig.Set("api_key", "mock-api-key")

// 		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer([]byte("invalid body")))
// 		w, ctx := setupTestContext(req)

// 		controller := &FavsController{}
// 		controller.Init(ctx, "FavsController", "Post", nil)
// 		controller.Post()

// 		assert.Contains(t, w.Body.String(), "Invalid JSON")
// 	})
// }
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/beego/beego/v2/server/web"
	// "github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

// func setupTestContext(req *http.Request) (*httptest.ResponseRecorder, *context.Context) {
// 	w := httptest.NewRecorder()
// 	ctx := &context.Context{
// 		Input:  &context.BeegoInput{Context: &context.Context{}},
// 		Output: &context.BeegoOutput{Context: &context.Context{}},
// 	}
// 	ctx.Reset(w, req)
// 	return w, ctx
// }

func TestFavsController_Get(t *testing.T) {
	t.Run("Successful Response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "mock-api-key", r.Header.Get("x-api-key"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":1,"image":{"id":"cat1","url":"http://example.com/cat1.jpg"}}]`))
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/favs", nil)
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Get", nil)
		controller.Get()

		assert.NotEmpty(t, w.Body.String())
		assert.Contains(t, w.Body.String(), "http://example.com/cat1.jpg")
	})

	t.Run("Missing API Key Configuration", func(t *testing.T) {
		web.AppConfig.Set("api_key", "")

		req, _ := http.NewRequest("GET", "/favs", nil)
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Get", nil)
		controller.Get()

		assert.NotEmpty(t, w.Body.String())
		assert.Contains(t, w.Body.String(), "Failed to fetch API key from configuration")
	})

	t.Run("Error in Fetching Favorites", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Internal Server Error"}`))
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/favs", nil)
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Get", nil)
		controller.Get()

		assert.NotEmpty(t, w.Body.String())
		assert.Contains(t, w.Body.String(), "Error fetching favorites")
	})

	t.Run("Timeout Scenario", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(6 * time.Second)
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		req, _ := http.NewRequest("GET", "/favs", nil)
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Get", nil)
		controller.Get()

		assert.NotEmpty(t, w.Body.String())
		assert.Contains(t, w.Body.String(), "Request timed out")
	})
}

func TestFavsController_Post(t *testing.T) {
	t.Run("Successful Post", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "mock-api-key", r.Header.Get("x-api-key"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"SUCCESS"}`))
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		payload := map[string]string{"image_id": "cat1"}
		payloadJSON, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer(payloadJSON))
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Post", nil)
		controller.Post()

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "SUCCESS", response["message"])
	})

	t.Run("Missing API Key Configuration", func(t *testing.T) {
		web.AppConfig.Set("api_key", "")

		payload := map[string]string{"image_id": "cat1"}
		payloadJSON, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer(payloadJSON))
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Post", nil)
		controller.Post()

		assert.Contains(t, w.Body.String(), "Failed to fetch API key from configuration")
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		web.AppConfig.Set("api_key", "mock-api-key")

		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer([]byte("invalid body")))
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Post", nil)
		controller.Post()

		assert.Contains(t, w.Body.String(), "Invalid JSON")
	})

	t.Run("API Error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Failed to save favorite"}`))
		}))
		defer server.Close()

		web.AppConfig.Set("api_key", "mock-api-key")
		web.AppConfig.Set("api_base_url", server.URL)

		payload := map[string]string{"image_id": "cat1"}
		payloadJSON, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/favs", bytes.NewBuffer(payloadJSON))
		w, ctx := setupTestContext(req)

		controller := &FavsController{}
		controller.Init(ctx, "FavsController", "Post", nil)
		controller.Post()

		assert.Contains(t, w.Body.String(), "Failed to save favorite")
	})
}
