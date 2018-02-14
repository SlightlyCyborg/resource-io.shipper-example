package main

import (
	"golang.org/x/net/context"
	pb "resource-io/shipper/vessel-service/proto/vessel"
)

// Our gRPC service handler
type service struct {
	repo Repository
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
