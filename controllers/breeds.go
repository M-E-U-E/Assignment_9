package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	web.Controller
}

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}

func (c *BreedsController) Get() {
	// Fetch the API key and base URL from the app configuration
	apiKey, err := web.AppConfig.String("api_key")
	if err != nil {
		fmt.Println("Error fetching API key from configuration:", err)
		c.Data["json"] = map[string]string{"error": "Failed to fetch API key from configuration"}
		c.ServeJSON()
		return
	}

	apiBaseURL, err := web.AppConfig.String("api_base_url")
	if err != nil {
		fmt.Println("Error fetching API base URL from configuration:", err)
		c.Data["json"] = map[string]string{"error": "Failed to fetch API base URL from configuration"}
		c.ServeJSON()
		return
	}

	// Create a channel for API responses and errors
	breedChan := make(chan []Breed)
	errChan := make(chan error)

	go func() {
		// API URL to fetch breeds
		url := apiBaseURL + "/breeds"
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			errChan <- err
			return
		}
		req.Header.Add("x-api-key", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		var breeds []Breed
		if err := json.Unmarshal(body, &breeds); err != nil {
			errChan <- err
			return
		}

		breedChan <- breeds
	}()

	// Wait for the response or error
	select {
	case breeds := <-breedChan:
		// Pass data to the template
		c.Data["Breeds"] = breeds
	case err := <-errChan:
		fmt.Println("Error fetching breeds:", err)
		c.Data["json"] = map[string]string{"error": "Failed to fetch breed data"}
		c.ServeJSON()
		return
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		fmt.Println("Request timed out while fetching breed data")
		c.Data["json"] = map[string]string{"error": "Request timed out"}
		c.ServeJSON()
		return
	}

	c.TplName = "breeds.tpl"
}
