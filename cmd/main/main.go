package main

import (
	"database/sql"
	"fmt"
	"gobnb/server"
	"log"
	_ "github.com/lib/pq"
)


func main() {
	connQuery := fmt.Sprintf("postgresql://gobnb:gobnb@localhost/gobnb?sslmode=disable")
	db, err := sql.Open("postgres", connQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := server.NewServer(":8888")
	if err := s.StartAndServe(); err != nil {
		log.Fatal(err)
	}
	db.Close()

}
