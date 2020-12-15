## Overview

Data logging with Golang. A protocol buffer implementation for sanitizing log messages based on
Medium [article](https://medium.com/pipedrive-engineering/data-logging-with-golang-how-to-store-customer-details-securely-14d49f2cf992)

## Requirements

- [protobuf](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)
- [protoc-gen-go](go install google.golang.org/protobuf/cmd/protoc-gen-go)

### Usage

commands: protoc --go_out=./protocol ./protocol/*.proto

### Sources

- https://blog.golang.org/protobuf-apiv2
- https://developers.google.com/protocol-buffers/docs/reference/go-generated

