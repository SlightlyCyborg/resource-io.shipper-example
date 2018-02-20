package main

import (

	// Import the generated protobuf code
	"fmt"
	"log"

	"golang.org/x/net/context"

	userService "github.com/SlightlyCyborg/resource-io.shipper-example/user-service/auth"
	vesselProto "github.com/SlightlyCyborg/resource-io.shipper-example/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	pb "resource-io/shipper/consignment-service/proto/consignment"
)

func main() {

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	repo := &InMemConsignmentRepository{}
	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func auth(token string) error {
	// Auth here
	authClient := userService.NewAuthClient("go.micro.srv.auth", client.DefaultClient)
	authResp, err := authClient.ValidateToken(context.TODO(), &userService.Token{
		Token: token,
	})
	log.Println("Auth resp:", authResp)
	log.Println("Err:", err)
	if err != nil {
		return err
	}

	return nil
}
