package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	web.Controller
}

type VotingCat struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Fetches random cat images and renders the voting page
func (c *VotingController) Get() {
	// Fetch the API key and base URL from the app configuration
	apiKey, err := web.AppConfig.String("api_key")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch API key from configuration"}
		c.ServeJSON()
		return
	}

	apiBaseURL, err := web.AppConfig.String("api_base_url")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch API base URL from configuration"}
		c.ServeJSON()
		return
	}

	// Create a channel to receive the API response
	catChan := make(chan []VotingCat)
	errChan := make(chan error)

	// Fetch random cat image in a separate goroutine
	go func() {
		url := apiBaseURL + "/images/search?limit=1"
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

		var cats []VotingCat
		if err := json.Unmarshal(body, &cats); err != nil {
			errChan <- err
			return
		}

		catChan <- cats
	}()

	// Wait for the response or error
	select {
	case cats := <-catChan:
		// Check if any cat data was returned
		if len(cats) > 0 {
			c.Data["CatImageURL"] = cats[0].URL
		} else {
			c.Data["CatImageURL"] = "/static/images/cat-placeholder.jpg" // Fallback image
		}
	case err := <-errChan:
		c.Data["json"] = map[string]string{"error": "Failed to fetch data: " + err.Error()}
		c.ServeJSON()
		return
	case <-time.After(5 * time.Second): 
		c.Data["json"] = map[string]string{"error": "Request timed out"}
		c.ServeJSON()
		return
	}

	// Render the voting page
	c.Data["UnixTimestamp"] = time.Now().Unix()
	c.TplName = "voting.tpl"
}
