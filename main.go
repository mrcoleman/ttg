package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrcoleman/ttg/ttgDB"
)

func Get_Home(db *sql.DB, r render.Render) {
	r.HTML(200, "home/index",nil)
}

func main() {
	m := martini.Classic()
	m.Map(ttgDB.SetupDB())
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".html", ".tmpl"},
		Charset:    "UTF-8",
	}))
	m.Get("/", Get_Home)
	m.Run()
}
