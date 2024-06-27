package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
)

type Client struct {
	Db  *sql.DB
	Dot *dotsql.DotSql
}

func NewClient() (*Client, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbName, host, port)
	client, err := sql.Open("postgres", connStr)
	checkErr(err)

	// Load queries
	dot, err := dotsql.LoadFromFile("./db/sql/locations/queries.sql")

	checkErr(err)

	return &Client{Db: client, Dot: dot}, nil
}

func (client *Client) Close() {
	client.Db.Close()
}

func (client *Client) InitTables() error {
	log.Info("Creating tables")
	res, err := client.Dot.Exec(client.Db, "create-locations-table")

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(res)
	return nil
}

func Query[T any](dot *dotsql.DotSql, db *sql.DB, query string) ([]T, error) {
	res, err := dot.Query(db, query)
	checkErr(err)

	var result []T

	for res.Next() {
		var resultLocal T
		if err := sqlx.StructScan(res, &result); err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, resultLocal)
	}

	return result, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
