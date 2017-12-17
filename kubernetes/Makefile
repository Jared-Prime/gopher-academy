APP?=kubernetes-lesson
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH=amd64

clean:
	rm -f ${APP}
	GOPATH=$(shell pwd) go fmt handlers

build: clean
	GOPATH=${GOPATH}:$(shell pwd) GOOS=${GOOS} GOARCH=${GOARCH} \
	  go build -o ${APP} -ldflags \
	   "-s -w -X handlers.Release=${RELEASE} \
		  -X handlers.Commit=${COMMIT} \
		  -X handlers.BuildTime=${BUILD_TIME}"

container: build
	docker build -t jprime/${APP}:${RELEASE} .

run: container
	docker stop jprime/${APP}:${RELEASE} || true && docker rm jprime/${APP}:${RELEASE} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		jprime/${APP}:${RELEASE}

test:
	GOPATH=${GOPATH}:$(shell pwd) go test -v -race handlers