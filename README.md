# WeatherFetcher
WeatherFetcher Microservice receives data from OpenWeather API for a specified city.

## Getting Started
 1. Install GO: https://go.dev/
 2. Install Protocol buffer compiler, protoc, version 3: https://grpc.io/docs/protoc-installation/
 3. Install the protocol compiler plugins for Go using the following commands:

    - ``` $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 ```
    - ``` $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 ```

4. Update your PATH so that the protoc compiler can find the plugins:

    - ``` $ export PATH="$PATH:$(go env GOPATH)/bin" ```

## Running the Code
1. In a terminal navigate to the WeatherFetcher directory and run:
    - ``` $ go run weather_server/main.go ```

2. In a second terminal from the same directory run:
    - ``` $ go run weather_client/main.go ```
    - OR ``` $ go run weather_client/main.go --city <city_name> ```

    - if no city_name is defined it will default to "London"

3. To close, ctrl + c in the weather_server terminal
