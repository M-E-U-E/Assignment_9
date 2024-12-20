package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type FavsController struct {
	web.Controller
}

type Favorite struct {
    ID    string `json:"id"`
    Image struct {
        URL string `json:"url"`
    } `json:"image"`
}

func (c *FavsController) Get() {
	apiKey := "live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg"
	url := "https://api.thecatapi.com/v1/favourites"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch data"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var favs []Favorite
    json.NewDecoder(resp.Body).Decode(&favs)

    c.Data["json"] = favs
    c.ServeJSON()
}

