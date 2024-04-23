package main

import (
	"context"
	"errors"
	"log"
	"strings"

	svc "github.com/onedaydev/mygolang/grpc/multiple-services/service"
	users "github.com/onedaydev/mygolang/grpc/user-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userService struct {
	svc.UnimplementedUserServer
}

type repoService struct {
	svc.UnimplementedRepoServer
}

func (s *userService) GetUser(
	ctx context.Context,
	in *users.UserGetRequest,
) (*users.UserGetReply, error) {
	log.Printf("received request for user with Email: %s Id: %s\n",
		in.Email,
		in.Id,
	)
	components := strings.Split(in.Email, "@")
	if len(components) != 2 {
		return nil, errors.New("invalid email address")
	}
	u := users.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       29,
	}
	return &users.UserGetReply{User: &u}, nil
}

func registerServices(s *grpc.Server) {
	svc.RegisterUsersServer(s, &userService{})
	svc.RegisterRepoServer(s, &repoService{})
	reflection.Register(s)
}
