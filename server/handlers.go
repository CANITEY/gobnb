package server

import (
	"gobnb/models"

	"github.com/labstack/echo/v4"
)

type toast struct {
	Msg string
	Status string
}


func (s *Server) authRoutes() {
	s.S.POST("/signup", func (c echo.Context) error {
		values, err := c.FormParams()
		if err != nil {
			return err
		}


		if password1, password2 := values.Get("password"), values.Get("password2"); password1 != password2 {
			data := struct{
				Msg string
				Status string
			}{
				"passwords doesn't match",
				"danger",
			}
			return c.Render(500, "toast", data)
		}

		user := models.NewFromUrlValues(values)
		if err := s.d.CreateUser(user); err != nil {
			data := toast{
				err.Error(),
				"danger",
			}
			return c.Render(500, "toast", data)
		}
		data := struct{
			Msg string
			Status string
		}{
			"success",
			"success",
		}
		return c.Render(200, "toast", data)
	})

	s.S.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")


		if email == "" || password == "" {
			data := toast{
				"email or password are empty",
				"danger",
			}
			return c.Render(403, "toast", data)
		}
		user, err := s.d.CheckUser(email, password)
		if err != nil {
			data := toast{
				err.Error(),
				"danger",
			}
			return c.Render(403, "toast", data)
		}

		if user == nil {
			data := toast{
				"user doesn't exist",
				"danger",
			}
			return c.Render(403, "toast", data)
		}

		if user.Email != email || user.Password != password {
			data := toast{
				"email or password is invalid",
				"danger",
			}
			return c.Render(403, "toast", data)
		}

		return c.Render(200, "profile", user)
	})
}
