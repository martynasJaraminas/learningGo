package main

import (
	"log"
	"net/http"
	"os"

	"parkingSharing/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Create location with parking lots
// POST location
// PUT location
// DELETE location
// GET location
// GET locations

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := db.NewClient()

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	client.InitTables()

	client.GetLocations()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	port := os.Getenv("SERVER_PORT")
	log.Print("Server is running on port: ", port)
	e.Logger.Fatal(e.Start(":" + port))
}
