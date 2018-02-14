package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	pb "resource-io/shipper/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "The Dagney Taggart", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	repo := &InMemVesselRepostiory{}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
