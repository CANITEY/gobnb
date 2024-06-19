package models

import "net/url"

type User struct {
	Id int `json:"id"`
    Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string
}

func NewFromUrlValues(values url.Values) User {
	return User{
		Name: values.Get("name"),
		Email: values.Get("email"),
		Password: values.Get("password"),
	}
}
