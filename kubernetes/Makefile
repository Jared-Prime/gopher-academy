APP?=kubernetes-lesson
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH=amd64

PROJECT?=github.com/jared-prime/gopher-academy/kubernetes/handlers

clean:
	rm -f ${APP}
	go fmt ${PROJECT}

build: clean
	GOOS=${GOOS} GOARCH=${GOARCH} \
	  go build -o ${APP} -ldflags \
	   "-s -w -X handlers.Release=${RELEASE} \
		  -X handlers.Commit=${COMMIT} \
		  -X handlers.BuildTime=${BUILD_TIME}"

container: build
	docker build -t jprime/${APP}:${RELEASE} .

push: container
	docker push jprime/${APP}:${RELEASE}

minikube: push
	for t in $(shell find . -type f -name "*.yaml"); do \
		cat $$t | \
			sed -E "s/\{\{ \.Release \}\}/\"$(RELEASE)\"/g" | \
			sed -E "s/\{\{ \.ServiceName \}\}/\"$(APP)\"/g"; \
		echo ---; \
	done > tmp.yaml
	kubectl apply -f tmp.yaml

run: container
	docker stop jprime/${APP}:${RELEASE} || true && docker rm jprime/${APP}:${RELEASE} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		jprime/${APP}:${RELEASE}

test:
	go test -v -race ${PROJECT}