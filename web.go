package main

import (
	"github.com/go-martini/martini"
	"github.com/kirigayakazushin/mkacgAPI/modules"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))

	m.Get("/news", modules.News)
	m.Get("/bangumi", modules.News)
	m.Run()
}
