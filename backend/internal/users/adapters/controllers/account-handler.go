package controllers

import (
	"context"
	"github.com/bufbuild/connect-go"
	v1 "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1/usersv1connect"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
	"github.com/ghandic/grpc-web-go-react-example/backend/internal/users/domain"
	"log"
)

type UserHandler struct {
	usersv1connect.UnimplementedUserServiceHandler
	userUsecases UserUsecases
}

type UserUsecases interface {
	GetUser(ctx context.Context, UserId int32) (*users.User, error)
	CreateUser(ctx context.Context, Name string) (*users.User, error)
	ListUsers(ctx context.Context, Req *v1.ListUsersRequest) (*domain.ListUsersResponse, error)
	DeleteUser(ctx context.Context, UserId int32) (bool, error)
}

func NewUserHandler(userUsecases UserUsecases) *UserHandler {
	return &UserHandler{userUsecases: userUsecases}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	user, err := h.userUsecases.CreateUser(ctx, req.Msg.Name)
	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}
	v1User := &v1.User{
		Id:   user.ID,
		Name: user.Name,
	}

	return connect.NewResponse(&v1.CreateUserResponse{
		User: v1User,
	}), nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	user, err := h.userUsecases.GetUser(ctx, req.Msg.UserId)
	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	return connect.NewResponse(&v1.GetUserResponse{
		User: &v1.User{Id: user.ID, Name: user.Name},
	}), nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	successfulDelete, err := h.userUsecases.DeleteUser(ctx, req.Msg.UserId)
	if !successfulDelete || err != nil {
		log.Fatalf("Error occurred %v", err)
	}
	return connect.NewResponse(&v1.DeleteUserResponse{}), nil
}

func (h *UserHandler) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	listUsers, err := h.userUsecases.ListUsers(ctx, req.Msg)

	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	return connect.NewResponse(&v1.ListUsersResponse{Users: PGUsersToUsers(*listUsers.Users), Total: listUsers.TotalCount}), nil
}
