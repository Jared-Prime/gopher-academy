#!/usr/env ruby

require 'grpc'
require File.expand_path('protocol/healthcheck_pb', __dir__)
require File.expand_path('protocol/healthcheck_services_pb', __dir__)

sub = Healthcheck::Service::Stub.new('localhost:50051, :this_channel_is_insecure)

loop do
  puts stub.ok(Healthcheck::HealthcheckRequest.new).status
end
