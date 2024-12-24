package controllers

import (
    "github.com/beego/beego/v2/server/web"
)

type ConfigController struct {
    web.Controller
}

// GetAPIKey handles fetching the API key
func (c *ConfigController) GetAPIKey() {
    apiKey, err := web.AppConfig.String("api_key")
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to fetch API key"}
    } else {
        c.Data["json"] = map[string]string{"api_key": apiKey}
    }
    c.ServeJSON()
}
