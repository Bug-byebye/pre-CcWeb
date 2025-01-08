package main

import (
	"context"
	"log"
	"net/http"
//	"strings"

	pb "CC-web/grpc/weather_rpc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 前端地址
		//AllowOriginFunc: func(origin string) bool {
		//	return strings.HasPrefix(origin, "http://localhost:")
		//},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/weather", func(c *gin.Context) {

		log.Println("[INFO] Incoming request to /weather endpoint")

		city := c.Query("city")
		extensions := c.DefaultQuery("extensions", "base")
		log.Printf("[DEBUG] Query parameters - city: %s, extensions: %s", city, extensions)

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Printf("[ERROR] Failed to connect to gRPC service: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to weather service"})
			return
		}
		defer conn.Close()

		client := pb.NewWeatherServiceClient(conn)
		res, err := client.GetWeather(context.Background(), &pb.WeatherRequest{
			City:       city,
			Extensions: extensions,
		})

		if err != nil {
			log.Printf("[ERROR] Failed to fetch weather data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[INFO] Successfully fetched weather data for city: %s", city)
		log.Printf("[DEBUG] Weather response: %+v", res)

		c.JSON(http.StatusOK, gin.H{
			"city":        res.City,
			"weather":     res.Weather,
			"temperature": res.Temperature,
			"humidity":    res.Humidity,
			"wind":        res.Wind,
			"reporttime":  res.Reporttime,
		})
	})

	return router
}
