package database

import (
	"fmt"
	"gobnb/models"

	"github.com/go-passwd/validator"
)

func (d *DB) GetUserInfo(id int) (models.User, error) {
	row := d.QueryRow("SELECT id, name, email, phone from users where id=?", id)
	user := &models.User{}
	if err := row.Scan(user.Id, user.Name, user.Email, user.Phone); err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (d *DB) CreateUser(user models.User) error {

	if err := validator.New(
		validator.CommonPassword(fmt.Errorf("This password is common")),
		validator.ContainsAtLeast("*,.!'\"+-@#$%^&(){}/", 1, fmt.Errorf("Password lack of special characters add atleast one of these *,.!'\"+-@#$%%^&(){}/")),
		validator.MinLength(8, fmt.Errorf("Password is too short at least 8 digits")),
	).Validate(user.Password);

	err != nil {
		return err
	}

	if _, err := d.Exec("INSERT INTO users(name, email, phone, password) VALUES($1, $2, $3, $4)", user.Name, user.Email, user.Phone, user.Password);

	err != nil {
		return err
	}

	return nil
}
