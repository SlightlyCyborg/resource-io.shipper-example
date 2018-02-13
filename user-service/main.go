package main

import (
	"github.com/micro/go-micro"
	"log"
	pb "resource-io/user-service/auth"
)

func main() {
	repo := &InMemRepo{}

	tokenService := &TokenService{repo}

	// Create a new go-micro service struct.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		// COLLIN: I changed this service from auth to user.
		micro.Name("go.micro.srv.auth"),
	)

	// Init will parse the command line flags.
	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())

	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
