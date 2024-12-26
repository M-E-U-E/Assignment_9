package main

import (
	_ "cat_api/routers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"fmt"
)
func setupApplication() {
    web.SetStaticPath("/static", "static")
    web.InsertFilter("/static/*", web.BeforeStatic, func(ctx *context.Context) {
        ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
        ctx.Output.Header("Pragma", "no-cache")
        ctx.Output.Header("Expires", "0")
    })
    web.AddFuncMap("fav", func(input string) string {
        return fmt.Sprintf("Favorite ID: %s", input)
    })
}
func main() {
	// Routers for the application
	

	web.SetStaticPath("/static", "static")
	web.InsertFilter("/static/*", web.BeforeStatic, func(ctx *context.Context) {
		ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Output.Header("Pragma", "no-cache")
		ctx.Output.Header("Expires", "0")
	})
	web.AddFuncMap("fav", func(input string) string {
		return fmt.Sprintf("Favorite ID: %s", input)
	})
	setupApplication()
	web.Run()
}
