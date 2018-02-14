package main

import (
	"log"
	"os"

	pb "github.com/SlightlyCyborg/resource-io.shipper-example/vessel-service/proto/vessel"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewVesselServiceClient("go.micro.srv.vessel", microclient.DefaultClient)

	spec := &pb.Specification{Capacity: 40, MaxWeight: 50000}

	r, err := client.Create(context.TODO(), &pb.Vessel{
		Id:        "1",
		Capacity:  70,
		MaxWeight: 60000,
		Name:      "The Dagney Taggart",
		Available: true,
		OwnerId:   "1"})

	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.Vessel.Id)

	r, err = client.FindAvailable(context.Background(), spec)
	if err != nil {
		log.Fatalf("Could not complete request: %v", err)
	}

	if r.Vessel == nil {
		log.Fatalf("No vessel is available to carry your goods")
	} else {
		log.Printf("The vessel, %s, can take your consignment", r.Vessel.Name)
	}

	// let's just exit because
	os.Exit(0)
}
