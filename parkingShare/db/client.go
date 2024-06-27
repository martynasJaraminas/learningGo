package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
)

type Client struct {
	db *sql.DB
}

// TODO: later move in to locations package?
type Location struct {
	id         string
	name       string
	paringLots []string
	address    sql.NullString
}

func NewClient() (*Client, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbName, host, port)
	client, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Client{db: client}, nil
}

func (client *Client) Close() {
	client.db.Close()
}

func (client *Client) InitTables() error {
	log.Info("Creating tables")

	dot, err := dotsql.LoadFromFile("./db/queries.sql")
	checkErr(err)

	res, err := dot.Exec(client.db, "create-locations-table")

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(res)
	return nil
}

func (client *Client) GetLocations() []Location {
	dot, err := dotsql.LoadFromFile("./db/queries.sql")
	checkErr(err)

	rows, err := dot.Query(client.db, "get-locations")
	checkErr(err)
	defer rows.Close()

	var locations []Location

	for rows.Next() {
		var location Location
		// TODO: can this be a dynamic based on struct?
		if err := rows.Scan(&location.id, &location.name, pq.Array(&location.paringLots), &location.address); err != nil {
			log.Error(err)
			continue
		}
		locations = append(locations, location)
	}

	return locations
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
