package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/swithek/dotsqlx"
)

type Client struct {
	Db  *sqlx.DB
	Dot *dotsqlx.DotSqlx
}

func NewClient() *Client {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbName, host, port)
	client, err := sqlx.Connect("postgres", connStr)
	checkErr(err)

	// Load queries
	dot, err := dotsql.LoadFromFile("./db/sql/locations/queries.sql")
	checkErr(err)

	dotx := dotsqlx.Wrap(dot)

	return &Client{Db: client, Dot: dotx}
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

func Query[T any](dotx *dotsqlx.DotSqlx, db *sqlx.DB, query string) ([]T, error) {

	var result []T

	if err := dotx.Select(db, &result, query); err != nil {
		log.Error(err)
		return nil, err

	}

	return result, nil
}

func QuerySingle[T any](dot *dotsql.DotSql, db *sql.DB, query string, param string) (T, error) {
	res, err := dot.Query(db, query, param)
	checkErr(err)

	var result T

	if res.Next() {
		if err := sqlx.StructScan(res, &result); err != nil {
			log.Error(err)
			return result, err
		}
	}

	return result, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
