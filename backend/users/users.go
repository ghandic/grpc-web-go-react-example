package users

import (
	"context"

	"sync"

	"github.com/bufbuild/connect-go"

	v1 "github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1/usersv1connect"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	usersv1connect.UnimplementedUserServiceHandler
	mu    sync.Mutex
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
	u.mu.Lock()
	defer u.mu.Unlock()
	user := &v1.User{
		Id:   ksuid.New().String(),
		Name: req.Msg.Name,
	}
	if u.users == nil {
		u.users = map[string]*v1.User{}
	}
	u.users[user.Id] = user
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
