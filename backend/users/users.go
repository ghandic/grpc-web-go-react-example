package users

import (
	"context"
	"fmt"
	"os"

	"sync"

	"github.com/bufbuild/connect-go"
	pgx "github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1/usersv1connect"
	db "github.com/ghandic/grpc-web-go-react-example/backend/users/db"
)

type UserService struct {
	usersv1connect.UnimplementedUserServiceHandler
	mu    sync.Mutex
	Conn  *pgx.Conn
	users map[string]*v1.User
}

func (u *UserService) GetUser(
	ctx context.Context,
	req *connect.Request[v1.GetUserRequest],
) (*connect.Response[v1.GetUserResponse], error) {

	if req.Msg.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing user id")
	}
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[req.Msg.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return connect.NewResponse(&v1.GetUserResponse{
		User: user,
	}), nil
}

func (u *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[v1.CreateUserRequest],
) (*connect.Response[v1.CreateUserResponse], error) {

	if req.Msg.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing name")
	}

	q := db.New(u.Conn)

	pg_user, err := q.CreateUser(ctx, req.Msg.Name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
		return nil, status.Errorf(codes.Internal, "CreateUser failed: %v\n", err)
	}

	user := &v1.User{
		Id:   pg_user.Id,
		Name: pg_user.Name,
	}

	return connect.NewResponse(&v1.CreateUserResponse{
		User: user,
	}), nil
}

func (u *UserService) ListUsers(
	ctx context.Context,
	req *connect.Request[v1.ListUsersRequest],
) (*connect.Response[v1.ListUsersResponse], error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := []*v1.User{}
	for _, value := range u.users {
		users = append(users, value)
	}
	return connect.NewResponse(&v1.ListUsersResponse{Users: users}), nil
}

func (u *UserService) DeleteUser(
	ctx context.Context,
	req *connect.Request[v1.DeleteUserRequest],
) (*connect.Response[v1.DeleteUserResponse], error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.users[req.Msg.UserId]; ok {
		delete(u.users, req.Msg.UserId)
	}
	return connect.NewResponse(&v1.DeleteUserResponse{}), nil
}
