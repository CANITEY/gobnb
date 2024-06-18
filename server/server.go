package server

import (
	"fmt"
	"gobnb/database"
	"gobnb/models"
	"html/template"
	"io"
	"strconv"

	"github.com/MadAppGang/httplog/echolog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)


type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


type Server struct {
	S *echo.Echo
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
	s.S.Use(middleware...)
}

func (s *Server) init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/*.tmpl")),
	}
	s.S.Renderer = t
	s.S.Logger.SetLevel(log.INFO)
	s.UseMiddleWare(echolog.Logger())
	s.UseMiddleWare(middleware.Logger())
	s.UseMiddleWare(middleware.Recover())
	s.PublicRoutes()
}

func (s *Server) PublicRoutes() {
	s.S.GET("/", func(c echo.Context) error {
		return c.Render(200, "base", nil)
	})

	s.S.GET("/login", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	s.S.GET("/apartments", func(c echo.Context) error {
		query := c.QueryParam("q")
		apartments := []models.Apartment{}
		if query == "" {
			data, err := s.d.GetApartments()
			apartments = append(apartments, data...)
			if err != nil {
				return echo.NewHTTPError(500, fmt.Sprintf("error: %v", err))
			}
		} else {
			data, err := s.d.SearchApartment(query)
			apartments = append(apartments, data...)
			if err != nil {
				return echo.NewHTTPError(500, fmt.Sprintf("error: %v", err))
			}
		}

		data := struct {
		Apartments []models.Apartment
		}{
			Apartments: apartments,
		}
		fmt.Println(apartments)
		return c.Render(200, "apartments", data)
	})

	s.S.GET("/apartments/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		info, err := s.d.GetApartment(id)
		if err != nil {
			return err
		}
		return c.Render(200, "apartment", info)
	})
}

func (s *Server) StartAndServe() error {
	if err := s.d.Initialize(); err != nil {
		return err
	}
	s.init()
	return s.S.Start(s.address)
}
