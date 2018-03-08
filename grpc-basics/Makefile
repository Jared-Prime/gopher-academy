DOCKERHUB_USER?=jprime

.DEFAULT_GOAL := help

build-container: ## build a stadalone container image for your project
	docker build -t $(DOCKERHUB_USER)/grpc-basics .

gen-protocol: ## generate the protocol bindings
	protoc \
	--ruby_out=./src/ruby \
	--grpc_out=./src/ruby \
	--go_out=plugins=grpc:src/go \
	--plugin=protoc-gen-grpc=`which grpc_ruby_plugin` \
	./protocol/healthcheck.proto

run-container: ## run a standalone container for your project
	docker run -it --rm $(DOCKERHUB_USER)/grpc-basics

help: ## display help for available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
