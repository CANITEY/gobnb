package database

import (
	"gobnb/models"
)

func (d *DB) GetUserInfo(id int) (models.User, error) {
	row := d.QueryRow("SELECT id, name, email, phone from users where id=?", id)
	user := &models.User{}
	if err := row.Scan(user.Id, user.Name, user.Email, user.Phone); err != nil {
		return models.User{}, err
	}

	return *user, nil
}

