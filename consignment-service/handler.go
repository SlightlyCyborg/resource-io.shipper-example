package main

import (
	"log"

	vesselProto "github.com/SlightlyCyborg/resource-io.shipper-example/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	pb "resource-io/shipper/consignment-service/proto/consignment"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return s.repo
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.CreateRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	log.Println("Creating: ", req.Consignment)

	auth_err := auth(req.Token)
	if auth_err != nil {
		return auth_err
	}

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Consignment.Weight,
		Capacity:  int32(len(req.Consignment.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.Consignment.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	err = repo.Create(req.Consignment)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req.Consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	log.Println(req.Token)
	auth_err := auth(req.Token)
	if auth_err != nil {
		return auth_err
	}

	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
