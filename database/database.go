package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(username, password, hostname, databasename string) (*sql.DB, error) {
	connQuery := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", username, password, hostname, databasename)
	db, err := sql.Open("postgres", connQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) Initialize() (error) {
	if _, err := d.Exec(`CREATE TABLE IF NOT EXIST users(
	id SERIAL PRIMARY KEY,
	name varchar(25),
	email varchar(30),
	password varchar(30),
	phone varchar(15)
`); err != nil {
		return err
	}

	if _, err := d.Exec(`CREATE TABLE IF NOT EXIST apartments(
	id SERIAL PRIMARY KEY,
	title varchar(25),
	description varchar,
	price numeric(7, 2),
	rental bool,
	availabe bool
`); err != nil {
		return err
	}

	return nil
}
