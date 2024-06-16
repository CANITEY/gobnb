package database

import "gobnb/models"

func (d *DB) GetApartments() ([]models.Apartment, error) {
	rows, err := d.Query("SELECT id, title, description, price, rental, availabe")
	if err != nil {
		return nil, err
	}

	apartments := make([]models.Apartment, 0)

	for rows.Next() {
		apartment := new(models.Apartment)
		rows.Scan(
			apartment.Id,
			apartment.Title,
			apartment.Description,
			apartment.Price,
			apartment.Rental,
			apartment.Available,
		)

		apartments = append(apartments, *apartment)
	}

	return apartments, nil
}

func (d *DB) GetApartment(id int) (apartment models.Apartment, err error) {
	row := d.QueryRow("SELECT id, title, description, price, rental, availabe where id=?", id)
	err = row.Scan(
		&apartment.Id,
		&apartment.Title,
		&apartment.Description,
		&apartment.Price,
		&apartment.Rental,
		&apartment.Available,
	)
	return
}

func (d *DB) SearchApartment(term string) ([]models.Apartment, error) {
	rows, err := d.Query("SELECT id, title, description, price, rental, availabe where title like %?%", term)
	if err != nil {
		return nil, err
	}

	apartments := make([]models.Apartment, 0)

	for rows.Next() {
		apartment := new(models.Apartment)
		rows.Scan(
			apartment.Id,
			apartment.Title,
			apartment.Description,
			apartment.Price,
			apartment.Rental,
			apartment.Available,
		)

		apartments = append(apartments, *apartment)
	}

	return apartments, nil
}
