package server

import (
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
}

func NewServer(address string) *Server {
	return &Server{
		echo.New(),
		address,
	}
}

func (s *Server) Use(middleware ...echo.MiddlewareFunc) {
	s.s.Use(middleware...)
}

func (s *Server) Init() {
	s.s.GET("/", h echo.HandlerFunc, m ...echo.MiddlewareFunc)
}
