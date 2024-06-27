package main

import (
	"log"
	"net/http"
	"os"

	"parkingSharing/db"
	"parkingSharing/models/location"

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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/locations", func(c echo.Context) error {
		result, err := db.Query[location.Location](client.Dot, client.Db, "get-locations")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, result)
		// return c.JSON(http.StatusOK, client.GetLocations())
	})

	// e.GET("/locations/:id", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, client.GetLocation(id))

	// }

	port := os.Getenv("SERVER_PORT")
	log.Print("Server is running on port: ", port)
	e.Logger.Fatal(e.Start(":" + port))
}
