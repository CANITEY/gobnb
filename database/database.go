package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(username, password, hostname, databasename string) (*DB, error) {
	connQuery := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", username, password, hostname, databasename)
	db, err := sql.Open("postgres", connQuery)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (d *DB) Initialize() (error) {
	if _, err := d.Exec(`CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name varchar(25) NOT NULL,
	email varchar(30) UNIQUE NOT NULL,
	password varchar(30) UNIQUE NOT NULL,
	phone varchar(15)
)`); err != nil {
		return err
	}

	if _, err := d.Exec(`CREATE TABLE IF NOT EXISTS apartments(
	id SERIAL PRIMARY KEY,
	title varchar(25),
	address varchar,
	description varchar,
	price numeric(7, 2),
	rental bool,
	availabe bool
)`); err != nil {
		return err
	}

	return nil
}
