package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	pb "profobuf-example/gravatar"
)

const port = ":50051"

type gravatarService struct{}

func (s *gravatarService) Generate(ctx context.Context, in *pb.GravatarRequest) (*pb.GravatarResponse, error) {
	log.Printf("Received email %v with size %v", in.Email, in.Size)
	return &pb.GravatarResponse{Url: gravatar(in.Email, in.Size)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to listen on port!"))
	}

	server := grpc.NewServer()
	pb.RegisterGravatarServiceServer(server, &gravatarService{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start server!"))
	}
}

func gravatarHash(email string) [16]byte {
	return md5.Sum([]byte(email))
}

func gravatarURL(hash [16]byte, size uint32) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", hash, size)
}

func gravatar(email string, size uint32) string {
	hash := gravatarHash(email)

	return gravatarURL(hash, size)
}
