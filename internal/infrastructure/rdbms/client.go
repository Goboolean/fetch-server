package rdbms

import (
	"database/sql"
	"fmt"

	"github.com/Goboolean/common/pkg/resolver"
	_ "github.com/lib/pq"
)


type DB struct {
	Queries

	db *sql.DB
}

func NewDB(c *resolver.ConfigMap) (*DB, error) {

	user, err := c.GetStringKey("USER")
	if err != nil {
		return nil, err
	}

	password, err := c.GetStringKey("PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	database, err := c.GetStringKey("DATABASE")
	if err != nil {
		return nil, err
	}

	urlstring := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", urlstring)

	if err != nil {
		return nil, err
	}

	return &DB{
		Queries: *New(db),
		db:      db,
	}, nil
}


func (db *DB) Close() {
	db.db.Close()
}


func (db *DB) Ping() error {
	return db.db.Ping()
}