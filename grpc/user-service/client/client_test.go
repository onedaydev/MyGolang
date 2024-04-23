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

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

type dummyUserService struct {
	users.UnimplementedUsersServer
}

func (s *dummyUserService) GetUser(
	ctx context.Context,
	in *users.UserGetRequest,
) (*users.UserGetReply, error) {
	u := users.User{
		Id:        "test-123-a",
		FirstName: "john",
		LastName:  "doe",
		Age:       29,
	}
	return &users.UserGetReply{User: &u}, nil
}

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	users.RegisterUsersServer(s, &dummyUserService{})
	go func() {
		err := startServer(s, l)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return s, l
}

func TestGetUser(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := func(
		ctx context.Context, addr string,
	) (net.Conn, error) {
		return l.Dial()
	}

	conn, err := grpc.DialContext(
		context.Background(),
		"", grpc.WithInsecure(),
		grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatal(err)
	}

	c := getUserServiceClient(conn)
	result, err := getUser(
		c,
		&users.UserGetRequest{Email: "john@doe.com"},
	)
	if err != nil {
		t.Fatal(err)
	}

	if result.User.FirstName != "john" || result.User.LastName != "doe" {
		t.Fatalf("expected: john doe, Got: %s %s",
			result.User.FirstName,
			result.User.LastName,
		)
	}
}
