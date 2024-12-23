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
	apiKey := "live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg"
	url := "https://api.thecatapi.com/v1/images/search?limit=1"

	// Create a request to fetch cat image
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Add("x-api-key", apiKey)

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch data"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to read response"}
		c.ServeJSON()
		return
	}

	var cats []VotingCat
	if err := json.Unmarshal(body, &cats); err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to parse response"}
		c.ServeJSON()
		return
	}

	// Check if any cat data was returned
	if len(cats) > 0 {
		c.Data["CatImageURL"] = cats[0].URL
	} else {
		c.Data["CatImageURL"] = "/static/images/cat-placeholder.jpg" // Fallback image
	}

	// Render the voting page
	c.Data["UnixTimestamp"] = time.Now().Unix()
	c.TplName = "voting.tpl"
}
