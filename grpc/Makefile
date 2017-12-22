APP?=service

run: build
	./${APP}

build: clean
	go build ${APP}.go

clean:
	go fmt
	go fmt ./weather
	go vet ./wunderground
	rm -f ${APP}

test:
	go test -v -race ./weather
	go test -v -race ./wunderground