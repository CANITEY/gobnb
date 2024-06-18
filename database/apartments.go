package database

import "gobnb/models"

func (d *DB) GetApartments() ([]models.Apartment, error) {
	rows, err := d.Query("SELECT id, title, address, price, rental, availabe from apartments;")
	if err != nil {
		return nil, err
	}

	apartments := []models.Apartment{}

	for rows.Next() {
		apartment := new(models.Apartment)
		rows.Scan(
			&apartment.Id,
			&apartment.Title,
			&apartment.Address,
			&apartment.Price,
			&apartment.Rental,
			&apartment.Available,
		)

		apartments = append(apartments, *apartment)
	}

	return apartments, nil
}

func (d *DB) GetApartment(id int) (apartment models.Apartment, err error) {
	row := d.QueryRow("SELECT id, title, address, description, price, rental, availabe from apartments WHERE id=$1", id)
	err = row.Scan(
		&apartment.Id,
		&apartment.Title,
		&apartment.Address,
		&apartment.Description,
		&apartment.Price,
		&apartment.Rental,
		&apartment.Available,
	)
	return
}

func (d *DB) SearchApartment(term string) ([]models.Apartment, error) {
	rows, err := d.Query("SELECT id, title, address, price, rental, availabe from apartments where title like %?%", term)
	if err != nil {
		return nil, err
	}

	apartments := make([]models.Apartment, 0)

	for rows.Next() {
		apartment := new(models.Apartment)
		rows.Scan(
			apartment.Id,
			apartment.Title,
			apartment.Address,
			apartment.Price,
			apartment.Rental,
			apartment.Available,
		)

		apartments = append(apartments, *apartment)
	}

	return apartments, nil
}
