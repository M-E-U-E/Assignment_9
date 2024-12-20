package routers

import (
	"cat_api/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/voting", &controllers.VotingController{})
	web.Router("/breeds", &controllers.BreedsController{})
	web.Router("/favs", &controllers.FavsController{})
	web.Router("/api/favorites", &controllers.FavsController{}, "get:Get")

    beego.Router("/", &controllers.MainController{})
}
