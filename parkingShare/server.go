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

	client := db.NewClient()
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
	})

	// e.GET("/locations/:id", func(c echo.Context) error {
	// 	result, err := db.QuerySingle[location.Location](client.Dot, client.Db, "get-location-by-id", c.Param("id"))

	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, err)

	// 	}
	// 	return c.JSON(http.StatusOK, result)

	// })

	port := os.Getenv("SERVER_PORT")
	log.Print("Server is running on port: ", port)
	e.Logger.Fatal(e.Start(":" + port))
}
