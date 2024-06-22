package main

import (
	"context"
	"flag"
	"log"
	"time"

	wf "github.com/jmzfung/WeatherFetcher/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	city = flag.String("city", "London", "the city to get data for")
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect new client: %v", err)
	}
	defer conn.Close()

	wfclient := wf.NewWeatherFetcherClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	weatherData, err := wfclient.GetWeather(ctx, &wf.WeatherRequest{City: *city})
	if err != nil {
		log.Fatalf("Error during GetWeather: %v", err)
	}
	city := weatherData.GetCity()
	humidity := weatherData.GetHumidity()
	wind := weatherData.GetWindSpeed()
	temp := weatherData.GetTemperature()

	log.Printf("The current weather data for %s is:\n", city)
	log.Printf("Temperature: %.2f°C\n", temp)
	log.Printf("Humidity: %.2f%%\n", humidity)
	log.Printf("WindSpeed: %.2fm/s\n", wind)
}
