# grpc-with-go

This is the repository to provide an example code for the [blogpost](https://k.lelonek.me/grpc-in-go) about using gRPC with Go.

## Development

You may want to leverage [`Makefile`](./Makefile) to install [`protoc` compiler](https://manpages.debian.org/testing/protobuf-compiler/protoc.1.en.html). Tou can also use it to download [`protoc-gen-go`](https://github.com/golang/protobuf/tree/master/protoc-gen-go) and finally to compile [`gravatar.proto` file](gravatar/gravatar.proto).

To do all of these things at once, just run:

    make compile

Once finished, you should see [`gravatar.pb.go` file](gravatar/gravatar.pb.go) generated.

## Usage

Tu run the application, you will need to start the server (to handle requests) and a client (to receive responses). They can be executed correspondingly as:

    go run server/main.go

    go run client/main.go

and have to be run in separate terminal windows as the server is a long-running process, while a client is just a one-time call.
