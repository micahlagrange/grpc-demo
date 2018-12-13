# grpc-demo
Testing out grpc


# Install dependencies

Go to https://github.com/protocolbuffers/protobuf/blob/master/src/README.md

- Follow the instructions to download the protobuf binary for c++ or all
- Run `go get -u github.com/golang/protobuf/protoc-gen-go` to get the go generation


Then once you have that, clone this repo.

Then,

```shell
cd grpc-demo
mkdir protos
protoc -I=. --go_out=protos person.proto
```

`go run ./cli.go`

or `go install` and run from your `$GOBIN`