package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	apiKey := "live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg"
	url := "https://api.thecatapi.com/v1/favourites"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching favorites:", err)
		c.Data["FavoriteImages"] = []string{} // Fallback to empty list
		c.TplName = "favs.tpl"
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Parse the response
	var favorites []FavoriteCat
	if err := json.Unmarshal(body, &favorites); err != nil {
		fmt.Println("Error parsing favorites response:", err)
		c.Data["FavoriteImages"] = []string{}
		c.TplName = "favs.tpl"
		return
	}

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
	c.TplName = "favs.tpl"
}

// Add a new favorite
func (c *FavsController) Post() {
	apiKey := "live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg"
	url := "https://api.thecatapi.com/v1/favourites"

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
