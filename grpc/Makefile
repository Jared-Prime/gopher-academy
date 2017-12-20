APP?=forecast

run: build
	./forecast

build: clean
	go build forecast.go

clean:
	go fmt
	go fmt ./weather
	go vet ./weather/weather_underground_api
	rm -f ${APP}

test:
	go test -v -race ./weather
	go test -v -race ./weather/weather_underground_api