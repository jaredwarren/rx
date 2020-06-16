package main

import (
	"log"
	"net"

	"github.com/jaredwarren/rx/user-service/controllers"
	"github.com/jaredwarren/rx/user-service/proto/user"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	// setup db connection

	//

	//

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	// pb.RegisterUserServiceServer(s, &service{repo})
	user.RegisterUserServiceServer(s, &controllers.User{})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
