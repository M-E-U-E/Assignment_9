package controllers

import (
	"net/http"
	// "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/beego/beego/v2/server/web"
)

func TestConfigController_GetAPIKey_Valid(t *testing.T) {
	// Mock API key in configuration
	web.AppConfig.Set("api_key", "mock-api-key")

	// Create a mock request
	req, _ := http.NewRequest("GET", "/config/api-key", nil)

	// Set up the context and recorder
	w, ctx := setupTestContext(req)

	// Initialize and call the controller
	controller := &ConfigController{}
	controller.Init(ctx, "ConfigController", "GetAPIKey", nil)
	controller.GetAPIKey()

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "mock-api-key")
}

func TestConfigController_GetAPIKey_MissingAPIKey(t *testing.T) {
	// Clear the API key from the configuration to simulate a missing API key
	web.AppConfig.Set("api_key", "")

	// Create a mock request
	req, _ := http.NewRequest("GET", "/config/api-key", nil)

	// Set up the context and recorder
	w, ctx := setupTestContext(req)

	// Initialize and call the controller
	controller := &ConfigController{}
	controller.Init(ctx, "ConfigController", "GetAPIKey", nil)
	controller.GetAPIKey()

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code) // Beego always returns 200 for ServeJSON
	assert.Contains(t, w.Body.String(), `"error":"Failed to fetch API key"`)
}
