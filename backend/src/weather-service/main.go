package main

import (
	"log"
	"net"

	pb "CC-web/grpc/weather_rpc"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWeatherServiceServer(grpcServer, &WeatherService{})

	log.Println("Weather Service is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
