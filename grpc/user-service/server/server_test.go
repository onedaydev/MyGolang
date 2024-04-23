package main

import (
	"context"
	"log"
	"net"
	"testing"

	users "github.com/onedaydev/mygolang/grpc/user-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	registerServices(s)
	go func() {
		err := startServer(s, l)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return s, l
}

func TestUserService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := func(
		ctx context.Context, addr string,
	) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(),
		"", grpc.WithInsecure(),
		grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatal(err)
	}
	usersClient := users.NewUsersClient(client)
	resp, err := usersClient.GetUser(
		context.Background(),
		&users.UserGetRequest{
			Email: "john@doe.com",
			Id:    "foo-bar",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if resp.User.FirstName != "john" {
		t.Errorf(
			"expected FirstName to be: john, Got: %s",
			resp.User.FirstName,
		)
	}
}
