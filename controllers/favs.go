package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type FavsController struct {
	web.Controller
}

type FavoriteCat struct {
	ID    int    `json:"id"`
	Image struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"image"`
}

// Fetch favorites for display
func (c *FavsController) Get() {
	// Fetch the API key and base URL from the app configuration
	apiKey, err := web.AppConfig.String("api_key")
	if err != nil {
		c.Data["FavoriteImages"] = []string{} // Fallback to empty list
		c.TplName = "favs.tpl"
		fmt.Println("Error fetching API key from configuration:", err)
		return
	}

	apiBaseURL, err := web.AppConfig.String("api_base_url")
	if err != nil {
		c.Data["FavoriteImages"] = []string{}
		c.TplName = "favs.tpl"
		fmt.Println("Error fetching API base URL from configuration:", err)
		return
	}

	// Create a channel for API responses and errors
	favChan := make(chan []FavoriteCat)
	errChan := make(chan error)

	go func() {
		url := apiBaseURL + "/favourites"
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

		var favorites []FavoriteCat
		if err := json.Unmarshal(body, &favorites); err != nil {
			errChan <- err
			return
		}

		favChan <- favorites
	}()

	// Wait for the response or error
	select {
	case favorites := <-favChan:
		// Extract image URLs
		var favoriteImages []map[string]interface{}
		for _, fav := range favorites {
			if fav.Image.URL != "" {
				favoriteImages = append(favoriteImages, map[string]interface{}{
					"id":  fav.ID,
					"url": fav.Image.URL,
				})
			}
		}

		// Pass the data to the template
		c.Data["FavoriteImages"] = favoriteImages
	case err := <-errChan:
		fmt.Println("Error fetching favorites:", err)
		c.Data["FavoriteImages"] = []string{}
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		fmt.Println("Request timed out while fetching favorites")
		c.Data["FavoriteImages"] = []string{}
	}

	c.TplName = "favs.tpl"
}

// Add a new favorite
func (c *FavsController) Post() {
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

	url := apiBaseURL + "/favourites"

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.Data["json"] = map[string]string{"error": "Invalid request"}
		c.ServeJSON()
		return
	}

	var requestData map[string]string
	if err := json.Unmarshal(body, &requestData); err != nil {
		fmt.Println("Error parsing request data:", err)
		c.Data["json"] = map[string]string{"error": "Invalid JSON"}
		c.ServeJSON()
		return
	}

	imageID, exists := requestData["image_id"]
	if !exists {
		c.Data["json"] = map[string]string{"error": "Missing image_id"}
		c.ServeJSON()
		return
	}

	payload := map[string]string{"image_id": imageID}
	payloadJSON, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, ioutil.NopCloser(bytes.NewReader(payloadJSON)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Add("x-api-key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing API request:", err)
		c.Data["json"] = map[string]string{"error": "Failed to save favorite"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	c.Data["json"] = json.RawMessage(respBody)
	c.ServeJSON()
}
