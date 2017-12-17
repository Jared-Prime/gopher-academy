APP?=kubernetes
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ${APP}
	GOPATH=$(shell pwd) go fmt handlers
	GOPATH=$(shell pwd) go fmt version

build: clean
	GOPATH=${GOPATH}:$(shell pwd) go build -o ${APP} -ldflags \
	 "-s -w -X version.Release=${RELEASE} \
		-X version.Commit=${COMMIT} \
		-X version.BuildTime=${BUILD_TIME}"

run: build
	PORT=${PORT} ./${APP}

test:
	GOPATH=$(shell pwd) go test -v -race handlers version