[read the blog](blog.haiqus.com)

Disclosure: this is how *I* structure my projects utilizing [gRPC, the simple service definition framework](https://grpc.io/). It works well for me in building backend service projects, but as needs change so do implementation details.

## Protocol Definition

The first step will be to define the protocol. With gRPC, we do so using the protocol buffer syntax. I am assuming version 3 in this tutorial. Our service will simply perform a health-check; the client (or caller) will send a request and the server ( or receiver ) will always respond positively, if it can.

Here's the definition. Create a file called `healthcheck.proto`

```
syntax = "proto3";

service Healthcheck {
  rpc Ok(HealthcheckRequest) returns (HealthcheckResponse) {}
}

message HealthcheckRequest {}

message HealthcheckResponse {
  enum Status {
    Standby = 0;
    Starting = 1;
    Ready = 2;
  }
  Status status = 1;
}
```

This protocol says: a `Healthcheck` service must implement an rpc method called `Ok` that will receive a `HealthcheckRequest` message and return a `HealthcheckResponse` message. The `HealthcheckRequest` message has no additional fields or contents; the `HealthcheckResponse` message has one additional field containing an enum value (0, 1, or 2) which describes the state of the service.

This is a bare-bones protocol that can be used in composition with other protocols to expose "readiness" of a receiver to a caller that wishes to issue further service calls.

You can learn more about protocol buffers at [the developer guide](https://developers.google.com/protocol-buffers/docs/overview)

## Project Directory Setup

Next, I have opinions. As stated above, they work for me and your milage may vary based on your individual needs.

I tend to place my `*.proto` files under a folder named, very imaginatively, `protocol`. Generated language bindings go under `src/<language>/protocol`. Lastly, I tend to create a `Makefile` to script common commands and a `Dockerfile` to isolate the test/development environment.

Using [`tree`](http://mama.indstate.edu/users/ice/tree/) to display the directory structure, your project folder will eventually look something like this:

```
.
├── Dockerfile
├── Makefile
├── protocol
│   └── healthcheck.proto
├── run-example.sh
└── src
    ├── go
    │   ├── example.go
    │   └── protocol
    │       └── healthcheck.pb.go
    └── ruby
        ├── example.rb
        └── protocol
            ├── healthcheck_pb.rb
            └── healthcheck_services_pb.rb
```

## Writing into the Makefile

The gRPC command line tool can be admittedly unwieldy to use; that's why I prefer to script the commands I need with `make`. As such, we'll be able to quickly generate new versions of our protocol bindings with a much simpler command.

```
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
	./protocol/healthcheck.prot 

run-container: ## run a standalone container for your project
	docker run -it --rm $(DOCKERHUB_USER)/grpc-basics

help: ## display help for available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
```

## Writing into the Dockerfile

We'll just use one Dockerfile for now to keep things simple:

```
FROM ruby:2.5.0

WORKDIR /app

COPY src .

CMD ["bash"]
```

## Implement the Server

For fun, we'll write the server with go(lang). Make sure that your project directory has been setup in your `GOPATH`, otherwise you may find some issues when importing the generated protocol bindings. See https://golang.org/doc/code.html#GOPATH for details on proper setup.

Create a new file at `src/go/example.go` like so

```
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "./protocol"
)

type receiver struct{}

func (s *receiver) Ok(ctx context.Context, req *pb.HealthcheckRequest) (*pb.HealthcheckResponse, error) {
	return &pb.HealthcheckResponse{Status: 2}, nil // positive response, no errors
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterHealthcheckServer(s, &receiver{})
	reflection.Register(s)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
```

As implemented, this `receiver` will always return a `Ready` reponse according to our protocol definition. We can edit the logic as we wish; the takeaway concept here is that the logic of our receiver _implementation_ is a detail distinct from our protocol definition. 

## Implement the Client

We will implement our client in Ruby. We'll use the generate Ruby bindings to stub the connection. Create a file at `src/ruby/example.rb` as such:

```
#!/usr/env ruby

require 'grpc'
require File.expand_path('protocol/healthcheck_pb', __dir__)
require File.expand_path('protocol/healthcheck_services_pb', __dir__)

sub = Healthcheck::Service::Stub.new('localhost:50051, :this_channel_is_insecure)

loop do
  puts stub.ok(Healthcheck::HealthcheckRequest.new).status
end
```

Save the file and make it executable with `chmod +x src/ruby/example.rb`.

## Run example using Docker

Lastly, we'll create a shell script to run on starting the Docker container and amend the Dockerfile to accomodate our new server and client implementations.

The shell script:

```
#!/bin/sh

example &
./src/ruby/example.rb
```

And the Dockerfile, modified from above:

```
FROM golang:1.10 as builder

WORKDIR /builder

COPY src/go .

RUN go build example.go

FROM ruby:2.5.0

WORKDIR /app

COPY src .
COPY --from=builder /builder/example example

CMD ["run-example.sh"]
```

These changes will make use of [multistage builds](https://docs.docker.com/develop/develop-images/multistage-build/) so that we may compile the go binary without having to add the go compiler and dependencies to our application image.

## Conclusion

In a future post, I will write about some reflections I have about what it might mean to build a project with toolsets like gRPC and Docker in the 21st century. Until then, you can find this code at https://github.com/Jared-Prime/gopher-academy/grpc-basics. Please submit any problems or suggestions as issues via Github. Any other comments you may send to me directly at <a href="mailto:write@haiq.us">write@haiq.us</a>. Thanks for reading!