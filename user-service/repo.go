package main

import (
	"errors"
	pb "resource-io/shipper/user-service/auth"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

//For now, I am only implementing an in memory version of these calls. This is to help me get practice!

type InMemRepo struct {
	users []*pb.User
}

func (repo *InMemRepo) GetAll() ([]*pb.User, error) {
	return repo.users, nil
}

func (repo *InMemRepo) Get(id string) (*pb.User, error) {
	var selected_u *pb.User
	selected_u = nil

	for _, user := range repo.users {
		if selected_u.Id == id {
			selected_u = user
		}
	}

	if selected_u == nil {
		var err error
		err = errors.New("Couldn't find user with that id")
		return nil, err
	}

	return selected_u, nil
}

func (repo *InMemRepo) Create(user *pb.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *InMemRepo) GetByEmail(email string) (*pb.User, error) {
	var selected_u *pb.User
	selected_u = nil

	for _, user := range repo.users {
		if user.Email == email {
			//Does not return, this gets the last created entry with this email
			selected_u = user
		}
	}

	if selected_u == nil {
		var err error
		err = errors.New("Couldn't find user with that email")
		return nil, err
	}

	return selected_u, nil
}
