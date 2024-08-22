package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

type Booking struct {
	UserID            string `json:"user_id"`
	ChargingStationID string `json:"charging_station_id"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Home route
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Endpoint to get charging stations
	r.GET("/charging-stations", func(c *gin.Context) {
		location := c.Query("location") // e.g., "37.7749,-122.4194" (latitude,longitude)

		apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
		if apiKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
			return
		}

		client, err := maps.NewClient(maps.WithAPIKey(apiKey))
		if err != nil {
			log.Printf("Error creating Google Maps client: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Google Maps client"})
			return
		}

		latLng := maps.LatLng{
			Lat: 37.7749,
			Lng: -122.4194,
		}
		if location != "" {
			_, err = fmt.Sscanf(location, "%f,%f", &latLng.Lat, &latLng.Lng)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location format"})
				return
			}
		}

		request := &maps.NearbySearchRequest{
			Location: &latLng,
			Radius:   5000,
			Type:     "charging_station",
		}

		response, err := client.NearbySearch(context.Background(), request)
		if err != nil {
			log.Printf("Error searching for nearby charging stations: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve charging stations"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"stations": response.Results})
	})

	// Endpoint to book a slot
	r.POST("/book-slot", func(c *gin.Context) {
		var booking Booking
		if err := c.ShouldBindJSON(&booking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking data"})
			return
		}

		// For now, just log the booking data
		// In a real application, you would save this to a database
		log.Printf("Booking: %+v\n", booking)

		c.JSON(http.StatusOK, gin.H{"status": "Booking successful!"})
	})

	// Start the server
	r.Run(":8080")
}
