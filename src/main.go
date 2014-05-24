package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	// "html/template"
	// "log"
	// "net/http"
)

func restaurants(r render.Render) {
	r.HTML(200, "restaurant_picker/index", "Restaurant Picker") // hello is the template's name
}

func choose_town(r render.Render) {
	r.HTML(200, "restaurant_picker/choose_town", "") // hello is the template's name
}

func choose_restaurant(r render.Render) {
	r.HTML(200, "restaurant_picker/choose_restaurant", "") // hello is the template's name
}

func restaurant(r render.Render) {
	r.HTML(200, "restaurant_picker/restaurant", "") // hello is the template's name
}

func main() {
	m := martini.Classic()
	// m.Get("/", func() string {
	// 	return "Hello world!"
	// })

	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
	}))

	// m.Get("/demos/restaurant_picker", index)
	// m.Get("/demos/restaurant_picker/choose_town", choose_town)

	m.Group("/demos/restaurant_picker", func(r martini.Router) {
		r.Get("", restaurants)
		r.Get("/choose_town", choose_town)
		r.Get("/choose_restaurant", choose_restaurant)
		r.Get("/restaurant", restaurant)
	})

	// This will set the Content-Type header to "application/json; charset=UTF-8"
	m.Get("/api", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	m.Run()
	// log.Fatal(http.ListenAndServe(":3000", m))
}
