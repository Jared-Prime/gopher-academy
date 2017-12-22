APP?=service
PROJECT?=weather
DEVELOPER?=jprime
VERSION?=0.0.1

GOOS?=linux
GOARCH?=amd64

WEATHER_UNDERGROUND_API_KEY?=fcfdaad82a7274ab

forecast:
	docker run --name weathercache --rm -d redis || true
	docker run --name weatherservice --rm --link "weathercache:redis.weather.svc" \
		-e WEATHER_UNDERGROUND_API_KEY=${WEATHER_UNDERGROUND_API_KEY} \
		-it ${DEVELOPER}/${PROJECT}:${VERSION} \
		/service -day=6

run: test clean build dockerize
	$(MAKE) forecast

build: clean
	GOOS=${GOOS} GOARCH=${GOARCH} go build ${APP}.go

clean:
	go fmt
	go fmt ./weather
	go vet ./wunderground
	rm -f ${APP}

test:
	go test -v -race ./weather
	go test -v -race ./wunderground

dockerize:
	docker build -t ${DEVELOPER}/${PROJECT}:${VERSION} .
