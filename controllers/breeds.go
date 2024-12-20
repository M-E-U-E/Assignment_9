package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	web.Controller
}

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Origin string `json:"origin"`
	Description string `json:"description"`
}

func (c *BreedsController) Get() {
	// The Cat API key
	apiKey := "live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg"
	url := "https://api.thecatapi.com/v1/breeds"

	// Create an HTTP request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", apiKey)

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch breed data"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Parse the JSON response
	var breeds []Breed
	if err := json.Unmarshal(body, &breeds); err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to parse breed data"}
		c.ServeJSON()
		return
	}

	// Pass data to the template
	c.Data["Breeds"] = breeds
	c.TplName = "breeds.tpl"
}
