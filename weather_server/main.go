package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	wf "github.com/jmzfung/WeatherFetcher/proto"
	"google.golang.org/grpc"
)

const (
	apiKey = "395eea175046bac258309f6cdfd24507"
)

var (
	port                           = flag.Int("port", 8080, "The server port")
	badResponse wf.WeatherResponse = wf.WeatherResponse{City: "", Temperature: 0, Humidity: 0, WindSpeed: 0}
)

type server struct {
	wf.UnimplementedWeatherFetcherServer
}

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

func (*server) GetWeather(ctx context.Context, in *wf.WeatherRequest) (*wf.WeatherResponse, error) {
	var url string = fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", in.GetCity(), apiKey)
	resp, err := http.Get(url)

	if err != nil {
		var msg string = fmt.Sprintf("error fetching weather data: %s", err)
		fmt.Println(errors.New(msg))
		return &badResponse, errors.New(msg)
	}
	if resp.StatusCode != 200 {
		var msg string = fmt.Sprintf("%s: %s", resp.Status, url)
		fmt.Println(errors.New(msg))
		return &badResponse, errors.New(msg)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		var msg string = fmt.Sprintf("error reading response body: %v", err)
		fmt.Println(errors.New(msg))
		return &badResponse, errors.New("error reading response body")
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("error unmarshaling weather data:", err)
		return &badResponse, errors.New("error unmarshaling weather data")
	}

	return &wf.WeatherResponse{City: in.GetCity(), Temperature: weatherData.Main.Temp, Humidity: weatherData.Main.Humidity, WindSpeed: weatherData.Wind.Speed}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	wf.RegisterWeatherFetcherServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
