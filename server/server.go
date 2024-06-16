package server

import (
	"html/template"
	"github.com/labstack/echo/v4"
	"github.com/MadAppGang/httplog/echolog"
	"io"
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
}

func NewServer(address string) *Server {
	return &Server{
		echo.New(),
		address,
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
		return c.Render(200, "apartments", nil)
	})
}

func (s *Server) StartAndServe() error {
	return s.s.Start(s.address)
}
