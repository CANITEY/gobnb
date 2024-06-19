package server

import (
	"gobnb/models"

	"github.com/labstack/echo/v4"
)



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
			return c.Render(200, "toast", data)
		}

		user := models.NewFromUrlValues(values)
		if err := s.d.CreateUser(user); err != nil {
			data := struct{
				Msg string
				Status string
			}{
				err.Error(),
				"danger",
			}
			return c.Render(200, "toast", data)
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
}
