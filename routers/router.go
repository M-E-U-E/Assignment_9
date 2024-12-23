package routers

import (
	"cat_api/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	// Enable CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "x-api-key", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	// Define routes
	beego.Router("/voting", &controllers.VotingController{})
	beego.Router("/breeds", &controllers.BreedsController{})
	beego.Router("/favs", &controllers.FavsController{})
	beego.Router("/api/favorites", &controllers.FavsController{}, "get:Get")
	beego.Router("/api/favorites", &controllers.FavsController{}, "post:Post")
	beego.Router("/", &controllers.MainController{})
}
