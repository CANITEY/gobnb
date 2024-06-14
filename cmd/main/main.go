package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/MadAppGang/httplog/echolog"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/*.tmpl")),
		}

	e := echo.New()
	e.Use(echolog.Logger())
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "base", "data")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(200, `
<title>HeHe</title>
<b>Hello world from golang</d>
`)
	})
	e.Start(":8888")
}
