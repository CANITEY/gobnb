package server

import (
	"database/sql"
	"gobnb/database"
	"html/template"
	"io"

	"github.com/MadAppGang/httplog/echolog"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


type Server struct {
	s *echo.Echo
	address string
	d *database.DB
}

func NewServer(address string) *Server {
	db, err := database.NewDB("gobnb", "gobnb", "localhost", "gobnb")
	if err != nil {
		panic(err.Error())
	}

	return &Server{
		echo.New(),
		address,
		db,
	}
}

func (s *Server) UseMiddleWare(middleware ...echo.MiddlewareFunc) {
	s.s.Use(middleware...)
}

func (s *Server) Init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/*.tmpl")),
	}
	s.s.Renderer = t
	s.UseMiddleWare(echolog.Logger())
	s.PublicRoutes()
}

func (s *Server) PublicRoutes() {
	s.s.GET("/", func(c echo.Context) error {
		return c.Render(200, "base", nil)
	})

	s.s.GET("/login", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	s.s.GET("/apartments", func(c echo.Context) error {
		if c.QueryParam("q") == "" {
			return c.Render(200, "apartments", nil)
		}
	})
}

func (s *Server) StartAndServe() error {
	return s.s.Start(s.address)
}
