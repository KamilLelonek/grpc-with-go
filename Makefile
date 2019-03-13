.PHONY: compile

PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

PROTOC := $(shell which protoc)
# If protoc isn't on the path, set it to a target that's never up to date, so the install command always runs.
ifeq ($(PROTOC),)
    PROTOC = must-rebuild
endif

# Figure out which machine you're running on.
UNAME := $(shell uname)

$(PROTOC):
# Run the right installation command for the operating system.
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
ifeq ($(UNAME), Linux)
	sudo apt-get install protobuf-compiler
endif

# Install $GOPATH/bin/protoc-gen-go unless exists.
$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go
	go install github.com/golang/protobuf/protoc-gen-go

compile: gravatar/gravatar.proto | $(PROTOC_GEN_GO) $(PROTOC)
	protoc --go_out=plugins=grpc:gravatar gravatar/*.proto --proto_path=gravatar
