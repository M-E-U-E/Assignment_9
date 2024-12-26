package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/beego/beego/v2/server/web"
)

// Import setupTestContext from test_helpers.go

func TestMainController_Get(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		// Create a mock request
		req, _ := http.NewRequest("GET", "/", nil)

		// Set up the context and recorder
		w, ctx := setupTestContext(req)

		// Initialize and call the controller
		controller := &MainController{}
		controller.Init(ctx, "MainController", "Get", nil)
		controller.Get()

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "beego.vip")
		assert.Contains(t, w.Body.String(), "astaxie@gmail.com")
	})

	t.Run("Client IP Test", func(t *testing.T) {
		// Mock a request with a specific client IP
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "127.0.0.1:12345"

		_, ctx := setupTestContext(req)

		controller := &MainController{}
		controller.Init(ctx, "MainController", "Get", nil)
		controller.Get()

		// Assert the client IP is included in the rendered data
		assert.Contains(t, controller.Data["ip"], "127.0.0.1")
	})

	t.Run("Template Name Test", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		_, ctx := setupTestContext(req)

		controller := &MainController{}
		controller.Init(ctx, "MainController", "Get", nil)
		controller.Get()

		// Assert the template name is correct
		assert.Equal(t, "index.tpl", controller.TplName)
	})
}
